package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)

	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			data := <-ch
			time.Sleep(1 * time.Millisecond)
			fmt.Println(data)
			wg.Done()
		}()
		ch <- 100
		wg.Wait()
	*/
}
