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
	timeoutCh := time.After(5 * time.Second)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-timeoutCh:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

//The below function can be replaced with time.After()
/*
func timeout(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
}
*/
