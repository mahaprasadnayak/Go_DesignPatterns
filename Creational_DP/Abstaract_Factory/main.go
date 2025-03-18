package main

import "fmt"

// Product Interface
type Shape interface {
	Draw()
}

// Concrete Products
type Circle struct{}

func (c Circle) Draw() {
	fmt.Println("Drawing a Circle")
}

type Square struct{}

func (s Square) Draw() {
	fmt.Println("Drawing a Square")
}

// Abstract Factory Interface
type ShapeFactory interface {
	CreateShape() Shape
}

// Concrete Factories
type CircleFactory struct{}

func (cf CircleFactory) CreateShape() Shape {
	return Circle{}
}

type SquareFactory struct{}

func (sf SquareFactory) CreateShape() Shape {
	return Square{}
}

// Client Code
func main() {
	var factory ShapeFactory

	// Use Circle Factory
	factory = CircleFactory{}
	shape1 := factory.CreateShape()
	shape1.Draw()

	// Use Square Factory
	factory = SquareFactory{}
	shape2 := factory.CreateShape()
	shape2.Draw()
}
