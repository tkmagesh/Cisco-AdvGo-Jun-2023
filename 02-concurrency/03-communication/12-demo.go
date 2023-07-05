package main

import (
	"fmt"
	"time"
)

// share memory by communicating

var result int

func main() {

	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("result :", result)

}

func add(x, y int, ch chan int) {
	result := x + y
	time.Sleep(5 * time.Second)
	ch <- result
}
