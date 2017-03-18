package client

import (
	"time"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

// Start cycles a signing day by day
func Start(path string) {
	log.Parse()
	go autoReply(path)
	for {
		log.Infofln("Signing begins at %v.", time.Now())
		c := new(config.C)
		c.Decode(path)
		bdussChan := make(chan string)
		go transBdussChan(bdussChan, c.BdussList)
		countMap := signByBDUSS(bdussChan, c)
		signingCount(countMap)
		countDown()
		broadcast()
		now := time.Now()
		sleepTime := time.NewTimer(tomorrow(now).Sub(now)).C
		log.Infofln("tomorrow's signing begins at %v.", tomorrow(now))
		<-sleepTime
	}
}

func autoReply(path string) {
	for {
		zone := time.FixedZone("BeiJing", 8*3600)
		now := time.Now().In(zone)
		start := today(now).Add(2 * time.Hour)
		end := start.Add(6 * time.Hour)
		if now.After(end) {
			start = start.AddDate(0, 0, 1)
			end = end.AddDate(0, 0, 1)
		}
		sleepTime := time.NewTimer(start.Sub(now)).C
		log.Infofln("Auto reply start at %s", start.Format(time.RFC3339))
		<-sleepTime
		c := new(config.C)
		c.Decode(path)
		if *debug {
			log.Infofln("config: %#v\n", c)
		}
		for _, bduss := range c.BdussList {
			getTopicList(bduss, c)
		}
		time.Sleep(time.Hour * 2)
	}
}
