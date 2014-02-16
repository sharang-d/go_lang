package main

import "fmt"

type Rectangle struct {
	height float32
	width  float32
}

func (r Rectangle) area() float32 {
	return r.height * r.width
}

func (r Rectangle) perimeter() float32 {
	return 2 * (r.height + r.width)
}


func main() {
	rect := Rectangle{10, 7}
	fmt.Println("Area =", rect.area())
	fmt.Println("Perimeter =", rect.perimeter())
}
