package main

import (
	"fmt"
	"time"
)

var intChan = make(chan int, 10)

func main() {
	go fmt.Println("a")
	go fmt.Println("b")
	go test()
	go CloseChan()
	time.Sleep(time.Second * 3)
	_, b := <-intChan
	fmt.Println(b)
	_, c := <-intChan
	fmt.Println(c)
}

func test() {
	time.Sleep(time.Second * 10)
	intChan <- 100
}

func CloseChan() {
	time.Sleep(time.Second * 20)
	close(intChan)
}
