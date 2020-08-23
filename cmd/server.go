package main

import (
	"dashboard/app"
	"net/http"
	_ "net/http/pprof"
	"os"

	log "github.com/sirupsen/logrus"
)

func initLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	level, err := log.ParseLevel(os.Getenv("APP_LOG_LEVEL"))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

func main() {
	if os.Getenv("ENABLE_PROFILING") == "true" {
		go func() {
			log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
		}()
	}

	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":9000"
	}

	initLogging()
	app.StartServer(addr)
}
