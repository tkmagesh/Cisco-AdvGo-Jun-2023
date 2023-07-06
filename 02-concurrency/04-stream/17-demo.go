package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genNos()
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		fmt.Println("All the values are received")
		break
	}
}

// producer
func genNos() <-chan int {
	ch := make(chan int)
	count := rand.Intn(10)
	fmt.Println("count = ", count)
	go func() {
		for i := 1; i <= count; i++ {
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
