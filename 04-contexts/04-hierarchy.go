package main

//modify the below program to generate the fibonocci numbers & prime numbers until the user hits the ENTER key
// make sure that generated values are printed in the main function as follows:
// fib : #
// prime : #

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()

	valCtx := context.WithValue(rootCtx, "root-key", "root-value")

	cancelCtx, cancel := context.WithCancel(valCtx)

	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := getGeneratedData(cancelCtx, wg)
	for data := range ch {
		fmt.Println(data)
	}
	wg.Wait()
}

func getGeneratedData(ctx context.Context, wg *sync.WaitGroup) <-chan string {
	fmt.Println("value of root-key in getGeneratedData : ", ctx.Value("root-key"))
	ch := make(chan string)
	dupValCtx := context.WithValue(ctx, "root-key", "new-root-value")
	ggDCtx := context.WithValue(dupValCtx, "ggd-key", "ggd-value")

	go func() {
		defer wg.Done()
		fibCtx, fibCancel := context.WithCancel(ggDCtx)
		defer fibCancel()
		wg.Add(1)
		fibCh := genFibonocci(fibCtx, wg)

		// primeCtx, primeCancel := context.WithCancel(ctx)
		primeCtx, primeCancel := context.WithTimeout(ggDCtx, 5*time.Second)
		defer primeCancel()
		wg.Add(1)
		primeCh := generatePrimes(primeCtx, wg)
	LOOP:
		for {
			select {
			case fibNo := <-fibCh:
				ch <- fmt.Sprintf("fib : %d", fibNo)
			case primeNo, isOpen := <-primeCh:
				if isOpen {
					ch <- fmt.Sprintf("prime : %d", primeNo)
				}
			case <-ctx.Done():
				fmt.Println("getGeneratedData - cancel signal received")
				close(ch)
				break LOOP
			}
		}
	}()
	return ch
}

func genFibonocci(ctx context.Context, wg *sync.WaitGroup) <-chan int {
	fmt.Println("value of root-key in genFibonocci : ", ctx.Value("root-key"))
	fmt.Println("value of ggd-key in genFibonocci : ", ctx.Value("ggd-key"))
	ch := make(chan int)
	go func() {
		defer wg.Done()
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("genFibonocci - cancel signal received")
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

func generatePrimes(ctx context.Context, wg *sync.WaitGroup) <-chan int {
	fmt.Println("value of root-key in generatePrimes : ", ctx.Value("root-key"))
	fmt.Println("value of ggd-key in generatePrimes : ", ctx.Value("ggd-key"))
	ch := make(chan int)
	go func() {
		defer wg.Done()
		no := 2
	LOOP:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("generatePrimes - cancel signal received")
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
