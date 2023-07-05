package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of f1
	f2()
	// time.Sleep(1 * time.Second) //blocking the execution of "main"
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
