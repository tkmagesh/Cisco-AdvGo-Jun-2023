package main

import (
	"fmt"
	"log"
	"time"
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

	/*
		logOperation(Add, 100, 200)
		logOperation(Subtract, 100, 200)
		logOperation(func(x, y int) {
			fmt.Println("Multiply Result:", x*y)
		}, 100, 200)
	*/

	/*
		logAdd := GetLogOperation(Add)
		logAdd(100, 200)

		logSubtract := GetLogOperation(Subtract)
		logSubtract(100, 200)
	*/

	/*
		profiledAdd := GetProfileOperation(Add)
		profiledAdd(100, 200)

		profiledSubtract := GetProfileOperation(Subtract)
		profiledSubtract(100, 200)
	*/

	logAdd := GetLogOperation(Add)
	profiledLogAdd := GetProfileOperation(logAdd)
	profiledLogAdd(100, 200)

	GetProfileOperation(GetLogOperation(Subtract))(100, 200)
	GetProfileOperation(GetLogOperation(func(x, y int) {
		fmt.Println("Multiply Result:", x*y)
	}))(100, 200)

	/*
		profileAdd := GetProfileOperation(Add)
		logProfileAdd := GetLogOperation(profileAdd)
		logProfileAdd(100, 200)
	*/
}

func GetProfileOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		elapsed := time.Since(start)
		log.Println("Operation took:", elapsed)
	}
}

func GetLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation Started")
		oper(x, y)
		log.Println("Operation Completed")
	}
}

// 3rd party library functions (cannot change this)
func Add(x, y int) {
	fmt.Println("Add Result:", x+y)
}

func Subtract(x, y int) {
	fmt.Println("Subtract Result:", x-y)
}
