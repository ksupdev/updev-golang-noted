package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("hi, This is demo Casttinge")
	// Convert int8 to int32
	var index int8 = 15
	bigIndex := int32(index)
	fmt.Printf("Convert int8 to int32 with int32(...) : %d \n", bigIndex)

	// Convert string to int
	lines_yesterday := "50"
	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Convert string to int with strconv.Atoi(...) : %d \n", yesterday)

	/*
		Cast interface to struct
		 - cast wiht "interface.(object)" =>  s.(rectangle)
	*/
	r := rectangle{width: 10, height: 10}
	c := circle{radius: 10}

	fmt.Printf("Rectangle area %f \n", getArea(r))
	fmt.Printf("Circle area %f \n", getArea(c))

	showInfo(r)
	showInfo(c)

	castToRectangle(r) // Casting Success
	castToRectangle(c) // Casting Error

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

func castToRectangle(s shape) {

	_, ok := s.(rectangle)
	if !ok {
		fmt.Println("Casting Error")
	} else {
		fmt.Println("Casting Success")
	}
}
