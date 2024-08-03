package behavior

import "fmt"

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) getType() string {
	return "Square"
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

type Circle struct {
	radius int
}

func (c *Circle) getType() string {
	return "Circle"
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) getType() string {
	return "Rectangle"
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(s *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) visitForRectangle(s *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func DoVisitor() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(areaCalculator)
	rectangle.accept(middleCoordinates)
}
