/*
Write a goroutine to generate prime numbers between the given "start" and "end"
consume the generate prime numbers and print them in the main function.
*/

package main

import "fmt"

func main() {
	ch := generatePrimes(2, 100)
	for prime := range ch {
		fmt.Println(prime)
	}
}

func generatePrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			if isPrime(i) {
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
