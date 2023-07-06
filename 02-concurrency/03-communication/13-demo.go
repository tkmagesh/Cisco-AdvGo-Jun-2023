package main

import (
	"fmt"
	"time"
)

// share memory by communicating
var result int

// consumer
func main() {
	go consumer()
}

func consumer() {
	ch := add(100, 200)
	/*
		go func() {
			ch <- 10000
		}()
	*/
	result := <-ch
	fmt.Println("result :", result)
}

// producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		time.Sleep(2 * time.Second)
		ch <- result
	}()
	return ch
}
