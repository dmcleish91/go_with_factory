package main

import "fmt"

type Shape interface {
	Draw() string
}

type Circle struct{}

type Rectangle struct{}

func (c Circle) Draw() string {
	return "Drawing a circle"
}

func (r Rectangle) Draw() string {
	return "Drawing a rectangle"
}

type ShapeFactory struct{}

func (sf ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return Circle{}
	case "rectangle":
		return Rectangle{}
	default:
		return nil
	}
}

/*
Suppose you're developing a web application that interacts with multiple databases, 
such as MySQL, PostgreSQL, and MongoDB. Each database requires a different type of 
connection object, and you want to abstract the process of creating these connections.
*/
func main() {
	factory := ShapeFactory{}

	circle := factory.CreateShape("circle")
	rectangle := factory.CreateShape("rectangle")

	fmt.Println(circle.Draw())
	fmt.Println(rectangle.Draw())
}