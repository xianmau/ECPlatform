package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"sync"
)

func main() {
	fmt.Println("tester runing...")
	t1 := time.Now().UnixNano()
	wait := new(sync.WaitGroup)
	for i := 0; i < 500; i ++{
		wait.Add(1)
		go func() {
			defer wait.Done()
			resp, err := http.Get("http://www.qq.com")
			if err != nil{
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//fmt.Println(err)
			}
			fmt.Println(len(body))
		}()
	}
	wait.Wait()
	t2 := time.Now().UnixNano()
	fmt.Println("total:", t2-t1)
}
