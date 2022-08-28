package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return r.width*2 + r.height*2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func measure (s shape) {
	fmt.Println("Type:", reflect.TypeOf(s), "\nArea:", s.area(), "\nPerimeter:", s.perimeter())
}

func main() {
	r := rectangle {width: 3, height: 4}
	c := circle {radius: 10}

	measure(r)
	measure(c)
}