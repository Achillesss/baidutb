package main

import (
	"flag"

	"github.com/achillesss/baidutb/client"
)

var path = flag.String("c", "./config/conf.toml", "config")

func main() {
	flag.Parse()
	client.Start(*path)
}
