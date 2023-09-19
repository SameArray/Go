package main

import "fmt"

func main() {
	rectangle := Rectangle{
		Width:  5,
		Height: 3,
	}

	f := 5
	fmt.Println(f)
	pointer := &f
	fmt.Println(pointer)
	fmt.Println(*pointer)

	rectangle.AppendToCalculatedAreas(45)
	fmt.Println(rectangle.CalculatedAreas)

	rectangle.EditFirsArea(90)
	fmt.Println(rectangle.CalculatedAreas)
}

type Rectangle struct {
	Width  int
	Height int

	CalculatedAreas []int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() int {
	return (r.Height + r.Width) * 2
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) AppendToCalculatedAreas(area int) {
	r.CalculatedAreas = append(r.CalculatedAreas, area)
}

func (r Rectangle) EditFirsArea(area int) {
	if len(r.CalculatedAreas) != 0 {
		r.CalculatedAreas[0] = area
	}
}