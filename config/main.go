package main

import (
	curr_dir "github.com/go-steven/curr-dir"
	"github.com/go-steven/goconfig/config"
)

func main() {
	cfg, err := config.ReadDefault(curr_dir.GetCurrDir() + "/config.cfg")
	if err != nil {
		panic(err.Error())
	}
	key, err := cfg.String("test", "key")
	if err != nil {
		panic(err.Error())
	}
	println("key=", key)
}
