package client

import (
	"flag"

	"github.com/achillesss/log"
)

var signingCountMap = make(map[string]int)
var debug = flag.Bool("debug", false, "debug mode")

func signingCount(countMap map[string]bool) {
	for k, v := range countMap {
		if v {
			signingCountMap[k]++
		}
	}
}

func broadcast() {
	for k, v := range signingCountMap {
		log.Infofln("signing count: %d\taccount: %q", v, k[:10])
	}
}
