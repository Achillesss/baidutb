package client

import (
	"time"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

func Start(path string) {
	for {
		c := new(config.C)
		c.Decode(path)
		var signTime *time.Time
		bdussChan := make(chan string)
		go transBdussChan(bdussChan, c.BdussList)
		for b := range bdussChan {
			go func(bduss string) {
				a := new(agent)
				a.configurate(c).checkConf().log()
				signTime = a.signOnePerson(bduss)
			}(b)
		}
		time.Sleep(time.Second)
		for i := 3; i > 0; i-- {
			log.Printfln("Signing in %d seoncd...", i)
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
