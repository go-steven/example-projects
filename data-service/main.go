package main

import (
	"github.com/bububa/goconfig/config"
	"github.com/bububa/mymysql/autorc"
	_ "github.com/bububa/mymysql/thrsafe"
	log "github.com/go-steven/logger"
	"time"
)

const (
	_CONFIG_FILE = "/var/code/go/config.cfg"
)

var (
	logger = log.NewLogger("")
)

func main() {
	cfg, _ := config.ReadDefault(_CONFIG_FILE)

	host, _ := cfg.String("masterdb", "host")
	user, _ := cfg.String("masterdb", "user")
	passwd, _ := cfg.String("masterdb", "passwd")
	dbname, _ := cfg.String("masterdb", "dbname")

	mdb := autorc.New("tcp", "", host, user, passwd, dbname)
	mdb.Register("set names utf8")

	userService := NewUsersService(mdb)
	go userService.Start()
	defer userService.Stop()

	time.Sleep(10 * time.Minute)
}
