package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		Add(100, 200)
		Subtract(100, 200)
	*/

	/*
		log.Println("Operation Started")
		Add(100, 200)
		log.Println("Operation Completed")

		log.Println("Operation Started")
		Subtract(100, 200)
		log.Println("Operation Completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	logOperation(Add, 100, 200)
	logOperation(Subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Multiply Result:", x*y)
	}, 100, 200)
}

func logOperation(oper func(int, int), x, y int) {
	log.Println("Operation Started")
	oper(x, y)
	log.Println("Operation Completed")
}

func logAdd(x, y int) {
	log.Println("Operation Started")
	Add(x, y)
	log.Println("Operation Completed")
}

func logSubtract(x, y int) {
	log.Println("Operation Started")
	Subtract(x, y)
	log.Println("Operation Completed")
}

// 3rd party library functions (cannot change this)
func Add(x, y int) {
	fmt.Println("Add Result:", x+y)
}

func Subtract(x, y int) {
	fmt.Println("Subtract Result:", x-y)
}
