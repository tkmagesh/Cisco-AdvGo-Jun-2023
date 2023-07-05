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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Length float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Width)
}

type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/*
func PrintStats(x interface {
	interface{ Area() float32 }
	interface{ Perimeter() float32 }
}) {
	PrintArea(x)      // x should be interface{ Area() float32 }
	PrintPerimeter(x) // x should be interface{ Perimeter() float32 }
}
*/

/*
func PrintStats(x interface {
	AreaFinder
	PerimeterFinder
}) {
	PrintArea(x)      // x should be AreaFinder
	PrintPerimeter(x) // x should be PerimeterFinder
}
*/

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintStats(x ShapeStatsFinder) {
	PrintArea(x)      // x should be AreaFinder
	PrintPerimeter(x) // x should be PerimeterFinder
}

func main() {
	c := Circle{5}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{10, 12}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

}
