package main

import (
	"fmt"
	"stress-test/request"
	"time"
)

func main() {
	//ch := make(chan int)
	//url := "http://www.baidu.com"
	//for i:=0; i<10; i++ {
	//	go request.Get(url, ch)
	//}
	//for i:=0; i<10; i++ {
	//	code := <- ch
	//	fmt.Println(code)
	//}
	//time.Sleep(1 * time.Second)
	fmt.Print(11111)
	fmt.Print(33333)
	fmt.Print(2222)
	fmt.Print(55555)
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go request.TestConnect(ch)
	}
	for i := 0; i < 10; i++ {
		res := <-ch
		fmt.Println(res)
	}
	time.Sleep(1 * time.Second)
}
