package handler

import (
	log "github.com/go-steven/logger"
	"github.com/kdar/factorlog"
)

var (
	Logger = log.NewLogger("")
)

func SetLogger(l *factorlog.FactorLog) {
	Logger = l
}
