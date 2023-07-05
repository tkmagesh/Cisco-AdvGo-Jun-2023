package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(16)

	count := flag.Int("count", 0, "number of goroutines")
	flag.Parse()
	fmt.Printf("Starting %d goroutines... Hit ENTER to start\n", *count)
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	for i := 1; i <= *count; i++ {
		wg.Add(1)    //increment the counter by 1
		go fn(i, wg) //scheduling the execution of fn
	}
	wg.Wait() //block until the wg counter to become zero
	fmt.Println("Hit ENTER to shutdown...")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() //decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
