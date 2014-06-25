package main

import (
	"log"
	"net/http"
	"routers"
	"utils/global"
	"utils/logger"
)

func main() {
	logger.Trace("server started.")
	logger.Trace(global.Config)
	logger.Trace(global.Sessions)

	routers.Register()
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func init() {
	// 在这里进行全局初始化
}
