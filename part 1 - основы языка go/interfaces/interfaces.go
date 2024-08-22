package main

import (
	"fmt"
	"math"
)

//type Shape interface {
//	Area() float32
//}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func (sq Square) Area() float32 {
	return sq.sideLength * sq.sideLength
}

func (sq Square) Perimeter() float32 {
	return 4 * sq.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

// main is the entry point of the program. It initializes a Square
// and Circle object with their respective dimensions. Then, it prints
// the area of each shape by calling the printShapeArea function.
func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)
	fmt.Println("*******************")
	printInterface(square)
	printInterface(circle)
	/// print string
	printInterface("it is a string")
	/// print number
	printInterface(500)
}

func printShapeArea(shape Shape) {
	fmt.Println("Area:", shape.Area())
	fmt.Println("Perimeter:", shape.Perimeter())
}

// empty interface

// all type all values are equals to empty interfaces

func printInterface(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("int", t)
	case bool:
		fmt.Println("bool", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("unknown type", t)
	}
	// change i to str type
	str, ok := i.(string)
	if ok != true {
		return
	} else {
		fmt.Println("!! converting complete !!")
		fmt.Println("after converting value:", str)
	}

}
