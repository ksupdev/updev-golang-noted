package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println("Hi, This is Demo interface")
	r := rectangle{width: 10, height: 10}
	showInfo(r)
	fmt.Printf("Rectangel value %f \n", getArea(r))

	c := circle{radius: 10}
	showInfo(c)
	fmt.Printf("Circlr value %f \n", getArea(c))

}

// Circle struct
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Rectangle struct
type rectangle struct {
	width, height float64
}

func (c rectangle) area() float64 {
	return c.height * c.width
}

// Interface  Shape

type shape interface {
	area() float64
}

func getArea(s shape) float64 {
	return s.area()
}

// Cast Interface
func showInfo(s shape) {
	t := reflect.TypeOf(s).Name()
	switch t {
	case "rectangle":
		r := s.(rectangle)
		fmt.Printf("Reactangle Width: %v Height: %v \n", r.width, r.height)
		break
	case "circle":
		c := s.(circle)
		fmt.Printf("Circle radius: %v \n", c.radius)
	}
}
