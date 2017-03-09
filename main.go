package main

import (
	"flag"

	"github.com/Achillesss/baidutb/client"
)

var path = flag.String("c", "./config/conf.toml", "config")

func main() {
	flag.Parse()
	client.Start(*path)
}
