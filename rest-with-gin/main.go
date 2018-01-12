package main

import (
	"flag"
	"fmt"
	//"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/go-steven/example-projects/rest-with-gin/handler"
	"github.com/go-steven/example-projects/rest-with-gin/router"
	log "github.com/go-steven/logger"
	"math/rand"
	"runtime"
	"time"
)

var (
	logFlag  = flag.String("log", "", "set log path")
	portFlag = flag.Int("port", 12345, "set port")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	logger := log.NewLogger(*logFlag)
	handler.SetLogger(logger)

	gin.SetMode(gin.ReleaseMode)
	r := router.NewRouter()
	logger.Infof("Rest API Server started at:0.0.0.0:%d", *portFlag)
	defer func() {
		logger.Infof("Rest API Server exit from:0.0.0.0:%d", *portFlag)
	}()
	//endless.ListenAndServe(fmt.Sprintf(":%d", *portFlag), r)
	r.Run(fmt.Sprintf(":%d", *portFlag))
}
