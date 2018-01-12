package main

import (
	log "github.com/go-steven/logger"
)

var (
	logger = log.NewLogger("")
)

func main() {
	logger.Infof("hello world.")
}
