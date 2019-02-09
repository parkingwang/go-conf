package main

import (
	"fmt"
	"log"
	"github.com/parkingwang/go-conf"
)

func main() {
	if config, err := cfg.LoadConfig("conf.d"); nil != err {
		log.Fatal(err)
	} else {
		config.ForEach(func(name string, value interface{}) {
			fmt.Printf("##-> %s : %v\n", name, value)
		})
	}
}
