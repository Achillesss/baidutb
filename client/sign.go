package client

import (
	"fmt"
	"time"

	"github.com/achillesss/baidutb/config"
)

func Start(path string) {
	c := new(config.C)
	c.Decode(path)
	a := new(agent)
	a.configurate(c).checkConf().log()
	// a.debug = true
	a.getList().parseListResp()
	for {
		var signTime *time.Time
		for k := range a.KwList {
			fmt.Printf("\n")
			a.Kw = k
			if a.err == nil {
				if a.getFid().parseFidResp().getTbs().parseTbsResp().canSign() {
					now := a.sign().signUp()
					if signTime == nil {
						signTime = &now
					}
					fmt.Printf("Time: %s\n", now.Format(time.RFC3339))
				}
			}
		}
		sleepTime := tomorrow(*signTime).Sub(*signTime)
		fmt.Printf("Today's signing ended at %s. tomorrow's signing begins at%s. sleep time: %v s", signTime.Format(time.RFC3339), tomorrow(*signTime), sleepTime.Seconds())
		a.getList().parseListResp()
		time.Sleep(sleepTime)
	}
}
