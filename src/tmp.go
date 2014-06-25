package main

import (
	"log"
	"time"
	"utils/tools"
)

func main() {
	log.Println(tools.GetDateYYYYMMDD())
	log.Println(tools.GetAppPath())

	time.Sleep(5 * time.Second)

}
