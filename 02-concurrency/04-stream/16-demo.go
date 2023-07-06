package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	count := 10
	ch := genNos(count)
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}

}

// producer
func genNos(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			ch <- i * 10
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return ch
}
