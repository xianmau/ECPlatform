package main

import (
	"log"
	"net/http"
	"routers"
	"utils/logger"
)

func main() {
	logger.Trace("server started.")
	routers.Register()
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
