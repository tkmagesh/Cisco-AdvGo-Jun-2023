package main

import "fmt"

type NumberType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Aute mollit dolore incididunt velit pariatur ex duis enim officia. Quis cillum reprehenderit non fugiat officia exercitation laboris sit magna. Laboris amet ipsum fugiat fugiat fugiat. Aliquip do irure ipsum id consectetur duis dolore officia ipsum. Cillum do est cillum labore Lorem ut proident quis sit sint esse sit nostrud irure. In culpa veniam do exercitation veniam incididunt do fugiat."
	x = true
	x = []int{1, 2, 3}
	x = struct{}{}
	x = 12.99

	// var x int
	// x = "100"
	x = 100
	// not run-time safe
	/*
		y := x.(int) + 200
		fmt.Println(y)
	*/

	// run-time safe (type assertion) (option - 1)
	if val, ok := x.(int); ok {
		fmt.Println("x is an int, x + 200 =", val+200)
	} else {
		fmt.Println("x is not an int")
	}

	// run-time safe (type assertion) (option -2)
	// x = "Eiusmod cillum sint pariatur est irure nisi commodo anim. Cillum qui officia fugiat magna cillum. Adipisicing ut in magna ullamco nisi deserunt labore reprehenderit non. Nulla exercitation minim et non cillum velit ut laborum eiusmod exercitation. Incididunt mollit dolore irure est nostrud ea incididunt aliquip consectetur eu. Tempor sit irure duis est pariatur aliquip consectetur. Nulla eu ex nulla ullamco nisi et occaecat aliquip dolor tempor."
	// x = true
	// x = struct{}{}
	x = []int{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is a struct, x =", val)
	default:
		fmt.Println("x is of an unknown type")
	}

	//type constraints
	printType(100)
	printType(100.99)

}

/*
func printType[T interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}](x T) {
	fmt.Println("x = ", x)
}
*/

func printType[T NumberType](x T) {
	fmt.Println("x = ", x)
}
