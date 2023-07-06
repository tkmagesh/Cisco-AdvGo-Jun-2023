package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to shutdown...")
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
	counter := 0
LOOP:
	for {
		select {
		case <-cancelCtx.Done():
			break LOOP
		default:
			counter++
			fmt.Println("counter :", counter)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
