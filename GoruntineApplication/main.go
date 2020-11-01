package main

import (
	"fmt"
	"time"
)

func main() {
	listChan := make(chan int, 10000)
	primeChan := make(chan int, 10000)
	exitChan := make(chan bool, 4)
	start := time.Now().Unix()
	var nummax int
	fmt.Scanln(&nummax)
	go putNumToChan(listChan, nummax)
	go putPrimeToChan(listChan, primeChan, exitChan)
	go putPrimeToChan(listChan, primeChan, exitChan)
	go putPrimeToChan(listChan, primeChan, exitChan)
	go putPrimeToChan(listChan, primeChan, exitChan)
	go func() {
		for i := 1; i <= 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()
	var i int = 1
	for {
		num, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(i, num)
		i++
	}
	end := time.Now().Unix()
	fmt.Println(end - start)
}

func putNumToChan(listChan chan<- int, nummax int) {
	for i := 2; i <= nummax; i++ {
		listChan <- i
	}
	close(listChan)
}

func putPrimeToChan(listChan chan int, primeChan chan int, exitChan chan bool) {
	for {
		var judge bool = true
		num, ok := <-listChan
		if !ok {
			exitChan <- true
			fmt.Println("取素数完毕。")
			break
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				judge = false
				break
			}
		}
		if judge {
			primeChan <- num
		}
	}
}
