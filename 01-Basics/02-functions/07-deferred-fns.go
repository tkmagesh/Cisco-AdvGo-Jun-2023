package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("	deferred [main]")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	// file = open file()
	defer func() {
		fmt.Println("	deferred [f1] - 1")
		// file.close()
	}()
	fmt.Println("f1 started")
	defer func() {
		fmt.Println("	deferred [f1] - 2")
	}()
	f2()
	/* .........  */
	fmt.Println("f1 completed")
}

func f2() {

	/*
		defer func() {
			fmt.Println("	deferred [f2] - 1")
		}()
	*/
	defer fmt.Println("	deferred [f2] - 1")
	fmt.Println("f2 started")
	/*
		defer func() {
			fmt.Println("	deferred [f2] - 2")
		}()
	*/
	defer fmt.Println("	deferred [f2] - 2")
	/* .........  */
	// os.Exit(0)
	fmt.Println("f2 completed")
}
