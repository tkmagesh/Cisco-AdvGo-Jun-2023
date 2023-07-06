package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}

}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * 10
			time.Sleep(2 * time.Second)
		}
	}()
	return ch
}
