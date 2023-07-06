package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFibonocci()
	for no := range ch {
		fmt.Println(no)
	}
}

// modify the below function to generate the values for a 5 second duration
func genFibonocci() <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < 10; i++ {
			ch <- x
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
		close(ch)
	}()
	return ch
}
