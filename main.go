package main

import (
	"fmt"
	"stress-test/request"
	"time"
)



func main() {
	ch := make(chan int)
	url := "http://www.baidu.com"
	for i:=0; i<10; i++ {
		go request.Get(url, ch)
	}
	for i:=0; i<10; i++ {
		code := <- ch
		fmt.Println(code)
	}
	time.Sleep(1 * time.Second)
}