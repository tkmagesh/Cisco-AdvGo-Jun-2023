package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

/*
// The following function has to be modified everytime a new shape is added
func PrintArea(x interface{}) {
	switch z := x.(type) {
	case Circle:
		fmt.Println("Area :", z.Area())
	case Rectangle:
		fmt.Println("Area :", z.Area())
	default:
		fmt.Println("Invalid type")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch z := x.(type) {
	case interface{ Area() float32 }:
		fmt.Println("Area :", z.Area())
	default:
		fmt.Println("Invalid type")
	}
}
*/

func PrintArea(x interface{ Area() float32 }) {
	fmt.Println("Area :", x.Area())
}

/*
type AreaFinder interface{ Area() float32 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}
*/

type Triangle struct {
}

func (t Triangle) Area() float32 {
	return 100
}

func main() {
	c := Circle{5}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	r := Rectangle{10, 12}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)

	t := Triangle{}
	PrintArea(t)

	PrintArea(100)

}
