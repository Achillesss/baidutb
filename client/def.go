package client

import (
	"time"

	"github.com/achillesss/log"
)

var signingTimeSlice []time.Time

var signingCountMap = make(map[string]int)

var debug bool

func pickSigningTime() (res time.Time) {
	for _, t := range signingTimeSlice {
		if t.After(res) {
			res = t
		}
	}
	return
}

func signingCount(countMap map[string]bool) {
	for k, v := range countMap {
		if v {
			signingCountMap[k]++
		}
	}
}

func broadcast() {
	log.Printfln("SIGNING STATISTICS:")
	for k, v := range signingCountMap {
		log.Printfln("signing count: %d\taccount: %q", v, k[:10])
	}
}
