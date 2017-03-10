package client

import (
	"time"

	"fmt"

	"github.com/achillesss/baidutb/config"
)

func tomorrow(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 500000000, now.Location())
}

func Start(path string) {
	c := new(config.C)
	c.Decode(path)
	a := new(agent)
	a.configurate(c).checkConf().log()
	a.debug = false
	var signTime time.Time
	for {
		for i, f := range c.FName {
			fmt.Printf("\n")
			a.Kw = f
			if a.err == nil {
				if a.getFid().parseFidResp().getTbs().parseTbsResp().canSign() {
					now := a.sign().signUp()
					if i == 0 {
						signTime = now
					}
					fmt.Printf("Time: %s\n", now.Format(time.RFC3339))
				}
			}
		}
		sleepTime := tomorrow(signTime).Sub(signTime)
		fmt.Printf("Today's signing ended at %s. tomorrow's signing begins at%s. sleep time: %v s", signTime.Format(time.RFC3339), tomorrow(signTime), sleepTime.Seconds())
		time.Sleep(sleepTime)
	}
}
