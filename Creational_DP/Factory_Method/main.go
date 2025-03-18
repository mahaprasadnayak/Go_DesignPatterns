package main

import "fmt"

type Triangle struct {
	base, height float64
}
type Rectangle struct {
	length, width float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
func (r *Rectangle) Area() float64 {
	return r.length * r.width
}
type Shape interface {
	Area() float64
}
// Factory Method for creating shapes
func ShapeFactory(shapeType string, dimensions ...float64) Shape {
	switch shapeType {
	case "triangle":
		if len(dimensions) < 2 {
			return nil // Not enough parameters
		}
		return &Triangle{base: dimensions[0], height: dimensions[1]}
	case "rectangle":
		if len(dimensions) < 2 {
			return nil // Not enough parameters
		}
		return &Rectangle{length: dimensions[0], width: dimensions[1]}
	default:
		return nil // Unknown shape type
	}
}


func main() {
	triangle := ShapeFactory("triangle", 10, 5)
	fmt.Println("Triangle Area:", triangle.Area())

	// Create a Rectangle using Factory Method
	rectangle := ShapeFactory("rectangle", 10, 5)
	fmt.Println("Rectangle Area:", rectangle.Area())
}