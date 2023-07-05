package main

import "fmt"

func main() {
	ch := make(chan int)

	/*
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	/*
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

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
