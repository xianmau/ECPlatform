package main

import (
	"net/http"
	"routers"
	"utils/global"
	log "utils/logger"
)

func main() {
	log.Trace("server started.")

	routers.Register()

	err := http.ListenAndServe(":"+global.Config.Get("port"), nil) //设置监听的端口
	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	// 在这里进行全局初始化
}
