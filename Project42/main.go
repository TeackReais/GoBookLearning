package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main")
	go func() {
		fmt.Println("Second Routine Sleep 1 Second")
		go func() {
			fmt.Println("Third Routine Sleep 2 Second")
			time.Sleep(2 * time.Second)
			fmt.Println("Live")
		}()
		time.Sleep(time.Second)
	}()
	time.Sleep(5 * time.Second)
}

/*协程之间关系平等，无父子关系*/
