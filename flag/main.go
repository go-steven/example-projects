package main

import (
	"flag"
)

var (
	portFlag = flag.Int("port", 12345, "set port")
)

func main() {
	flag.Parse()

	println("port:", *portFlag)
}
