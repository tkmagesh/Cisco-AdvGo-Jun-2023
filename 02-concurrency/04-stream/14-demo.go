package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		ch <- 10
		time.Sleep(2 * time.Second)
		ch <- 20
		time.Sleep(2 * time.Second)
		ch <- 30
		time.Sleep(2 * time.Second)
		ch <- 40
		time.Sleep(2 * time.Second)
		ch <- 50
		time.Sleep(2 * time.Second)
	}()
	return ch
}
