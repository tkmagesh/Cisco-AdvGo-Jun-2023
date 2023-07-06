package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	fmt.Printf("# of values in the channel : %d\n", len(ch))
	ch <- 20
	fmt.Printf("# of values in the channel : %d\n", len(ch))
	fmt.Println(<-ch)
	fmt.Printf("# of values in the channel : %d\n", len(ch))
	fmt.Println(<-ch)
	fmt.Printf("# of values in the channel : %d\n", len(ch))
}
