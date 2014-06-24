package main

import (
	"log"
	"net/http"
	"routers"
)

func main() {
	routers.Register()
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
