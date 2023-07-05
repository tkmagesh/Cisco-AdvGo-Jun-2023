package main

import "fmt"

func main() {

	exec(f1)
	exec(f2)
	exec(func() {
		fmt.Println("anonymous function invoked")
	})
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
