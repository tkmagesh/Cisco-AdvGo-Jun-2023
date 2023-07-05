package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrDivideByZeroError = errors.New("cannot divide by zero")

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("something went wrong! err :", err)
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	q, r, e := divideWrapper(100, divisor)
	if e != nil {
		fmt.Println("error:", e)
		return
	}
	fmt.Printf("dividing 100 by %d, quotient is %d, remainder is %d\n", divisor, q, r)
}

func divideWrapper(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party
func divide(x, y int) (int, int) {
	fmt.Println("calculating quotient")
	if y == 0 {
		panic(ErrDivideByZeroError)
	}
	quotient := x / y
	fmt.Println("calculating remainder")
	remainder := x % y
	return quotient, remainder
}
