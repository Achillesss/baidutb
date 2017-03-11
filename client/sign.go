package client

import (
	"fmt"
	"time"

	"github.com/achillesss/baidutb/config"
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
				fmt.Printf("\nSigning %s\n", bduss)
				signTime = a.signOnePerson(bduss)
			}(b)
		}
		time.Sleep(time.Second)
		for i := 3; i > 0; i-- {
			fmt.Printf("Signing in %d seoncd...\n", i)
			time.Sleep(time.Second)
		}
		if signTime == nil {
			t := time.Now()
			signTime = &t
		}
		sleepTime := tomorrow(*signTime).Sub(*signTime)
		fmt.Printf("Today's signing ended at %s. tomorrow's signing begins at%s. sleep time: %v s", signTime.Format(time.RFC3339), tomorrow(*signTime), sleepTime.Seconds())
		time.Sleep(sleepTime)
	}
}
