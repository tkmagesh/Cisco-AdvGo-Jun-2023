package main

//modify the below program to generate the values until the user hits the ENTER key
import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	ch := genFibonocci(stopCh)
	for no := range ch {
		fmt.Println(no)
	}
}

// modify the below function to generate the values for a 5 second duration
func genFibonocci(stopCh chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
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
