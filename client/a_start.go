package client

import (
	"time"

	"github.com/achillesss/baidutb/config"
	"github.com/achillesss/log"
)

// Start cycles a signing day by day
func Start(path string) {
	for {
		log.Parse()
		// pre signing
		// debug = true
		signingTimeSlice = nil
		c := new(config.C)
		c.Decode(path)

		// signing
		bdussChan := make(chan string)
		go transBdussChan(bdussChan, c.BdussList)
		countMap := signByBDUSS(bdussChan, c)
		signingCount(countMap)
		// post signing
		countDown()
		broadcast()
		s := pickSigningTime()
		sleepTime := tomorrow(s).Sub(s)
		log.Infofln("Today's signing ended at %s. tomorrow's signing begins at%s. sleep time: %v s", s.Format(time.RFC3339), tomorrow(s), sleepTime.Seconds())
		// sleep till next day's signing
		time.Sleep(sleepTime)
	}
}
