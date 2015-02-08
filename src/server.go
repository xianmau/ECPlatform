package main

import (
	"net/http"
	"routers"
	"utils/global"
	log "utils/glog"
	"flag"
	"runtime"
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	global.Init()
	routers.Register()
	log.Infof("Server started, listening port: %s\n", global.Config["port"])
	err := http.ListenAndServe(":"+global.Config["port"], nil)
	if err != nil {
		log.Fatalf("Server down: %s\n", err.Error())
	}
}
