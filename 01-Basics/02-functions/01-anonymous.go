package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi")
	}()

	func(name string) {
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	}("Magesh")

	var msg = func(name string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", name)
	}("Suresh")
	fmt.Print(msg)

	var q, r = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Println(q, r)

}
