package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	WithArea
	WithPerimeter
}

type WithArea interface {
	Area() float32
}

type WithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func NewSquare(sideLength float32) Square {
	return Square{
		sideLength: sideLength,
	}
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

func PrintShapeArea(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func PrintInterface(i interface{}) {
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
