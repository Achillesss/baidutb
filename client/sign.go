package client

import (
	"time"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

var timeChan chan *time.Time

func Start(path string) {
	for {
		c := new(config.C)
		c.Decode(path)
		var signTime *time.Time
		bdussChan := make(chan string)
		go transBdussChan(bdussChan, c.BdussList)
		timeChan = make(chan *time.Time)
		for b := range bdussChan {
			go func(bduss string) {
				a := new(agent)
				a.configurate(c).checkConf().log()
				a.signOnePerson(bduss)
			}(b)
		}

		select {
		case t := <-timeChan:
			signTime = t
		}

		time.Sleep(time.Second)
		for i := 3; i > 0; i-- {
			second := "seconds"
			if i == 1 {
				second = "second"
			}
			log.Printfln("Signing will end in %d %s...", i, second)
			time.Sleep(time.Second)
		}

		if signTime == nil {
			t := time.Now()
			signTime = &t
		}

		sleepTime := tomorrow(*signTime).Sub(*signTime)
		log.Printfln("Today's signing ended at %s. tomorrow's signing begins at%s. sleep time: %v s", signTime.Format(time.RFC3339), tomorrow(*signTime), sleepTime.Seconds())
		time.Sleep(sleepTime)
	}
}
