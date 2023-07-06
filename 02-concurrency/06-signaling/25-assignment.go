package main

//modify the below program to generate the fibonocci numbers & prime numbers until the user hits the ENTER key
// make sure that generated values are printed in the main function as follows:
// fib : #
// prime : #

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		close(stopCh)
	}()
	ch := getGeneratedData(stopCh)
	for data := range ch {
		fmt.Println(data)
	}
}

func getGeneratedData(stopCh <-chan struct{}) <-chan string {
	ch := make(chan string)
	go func() {
		fibCh := genFibonocci(stopCh)
		primeCh := generatePrimes(stopCh)
	LOOP:
		for {
			select {
			case fibNo := <-fibCh:
				ch <- fmt.Sprintf("fib : %d", fibNo)
			case primeNo := <-primeCh:
				ch <- fmt.Sprintf("prime : %d", primeNo)
			case <-stopCh:
				close(ch)
				break LOOP
			}
		}
	}()
	return ch
}

func genFibonocci(stopCh <-chan struct{}) <-chan int {
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

func generatePrimes(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		no := 2
	LOOP:
		for {
			select {
			case <-stopCh:
				break LOOP
			default:
				if isPrime(no) {
					ch <- no
					time.Sleep(300 * time.Millisecond)
				}
				no++
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
