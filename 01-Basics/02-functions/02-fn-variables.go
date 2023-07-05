package main

import "fmt"

func main() {
	/*
		var sayHi interface{}
		sayHi = func() {
			fmt.Println("Hi")
		}
		sayHi = 100
		if fn, ok := sayHi.(func()); ok {
			fn()
		} else {
			fmt.Println("sayHi is not a function")
		}
	*/
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()

	var printGreet func(string)
	printGreet = func(name string) {
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	}
	printGreet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(name string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", name)
	}
	msg := getGreetMsg("Suresh")
	fmt.Print(msg)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	var q, r = divide(100, 7)
	fmt.Println(q, r)

}
