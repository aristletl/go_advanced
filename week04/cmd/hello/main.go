package main

import (
	"flag"
	"week04/internal/conf"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "/Users/zhaotianlun/Project/go/go_advanced/week04/config/config.json", "config path, eg: -conf config.json")
}

func main() {
	flag.Parse()
	c := conf.New(flagconf)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc = &conf.Data{}
	if err := c.Scan(bc); err != nil {
		panic(err)
	}

	serve := initApp(bc)
	if err := serve.Run(); err != nil {
		panic(err)
	}
}
