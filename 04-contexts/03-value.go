package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	//root context
	rootCtx := context.Background()

	//context with data
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")

	//context with timeout
	cancelCtx, cancel := context.WithTimeout(valCtx, 5*time.Second)
	defer cancel()

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go fn(cancelCtx, wg)
	wg.Wait()
}

func fn(cancelCtx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("value of root-key in fn : ", cancelCtx.Value("root-key"))
	counter := 0
LOOP:
	for {
		select {
		case <-cancelCtx.Done():
			fmt.Println("done signal received")
			break LOOP
		default:
			counter++
			fmt.Println("counter :", counter)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
