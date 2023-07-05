package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	LogWrapper(Add)(100, 200)
	LogWrapper(Add2)(100, 200, 300)
}

func Add(x, y int) {
	fmt.Println("Add Result:", x+y)
}

func Add2(x, y, z int) {
	fmt.Println("Add2 Result:", x+y+z)
}

func LogWrapper(fn interface{}) func(args ...interface{}) {
	return func(args ...interface{}) {
		log.Println("Operation Started")
		fnValue := reflect.ValueOf(fn)
		fnType := fnValue.Type()

		inputs := make([]reflect.Value, len(args))
		for i, arg := range args {
			inputs[i] = reflect.ValueOf(arg)
		}

		fmt.Printf("Calling function %s with arguments %v\n", fnType, inputs)
		/*
			result := fnValue.Call(inputs)
			fmt.Printf("Result: %v\n", result)
		*/
		fnValue.Call(inputs)
		log.Println("Operation Completed")
	}
}
