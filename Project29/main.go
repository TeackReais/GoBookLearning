package main

import (
	"fmt"
)

func main() {
	StringChan := make(chan string, 50)
	IntChan := make(chan int, 50)
	for i := 1; i <= 50; i++ {
		StringChan <- fmt.Sprintf("String:%d", i)
	}
	for i := 1; i <= 50; i++ {
		IntChan <- i
	}
	for {
		select { //随机选择
		case v := <-StringChan:
			{
				fmt.Println(v)
				continue
			}
		case v := <-IntChan:
			{
				fmt.Println(v)
				continue
			}
		default:
			{
				fmt.Println("over")
				break
			}
		}
		break
	}
}
