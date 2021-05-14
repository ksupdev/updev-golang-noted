package main

//https://golangbot.com/structs-instead-of-classes/
import "fmt"

func main() {
	fmt.Println("Hi, this is demo Struct")
	var p1 product
	p1.name = "Andunio"
	p1.price = 100
	p1.stock = 20

	showProduct(p1)
	showProductWithPointer(&p1)
	updateWithPointer(&p1)

	// Show product after update
	showProduct(p1)

	//
	//p1 = p1.clear()

	fmt.Println("--- Show Struct function ---")
	fmt.Println("Before discount")
	showProduct(p1)
	p1 = p1.setDiscount(5) // past value and
	fmt.Println("After discount")
	showProduct(p1)
	p1 = p1.clear()
	fmt.Println("Clear")
	showProduct(p1)

}

type product struct {
	name  string
	price int
	stock int
}

// --- Defind product function ---
// '(p product)' use for defind Struct owner
func (p product) clear() product {
	p.price = 0
	p.stock = 0

	return p
}

func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}

// --- External func ---

func showProduct(p product) {
	fmt.Printf("[showProduct] : Display Product value = %v \n", p)
}

func showProductWithPointer(p *product) {
	fmt.Printf("[showProductWithPointer] : Display pointer value = %v \n", p)
	//[showProductWithPointer] : Display pointer value = &{Andunio 100 20}

	fmt.Printf("[showProductWithPointer] : Display real value = %v \n", *p)
	//[showProductWithPointer] : Display real value = {Andunio 100 20}

	/*
		&{Andunio 100 20}  <= display pointer
		if you want to print value , you should use add '*' *p
		{Andunio 100 20} <= display value
	*/
}

func updateWithPointer(p *product) {
	p.price = p.price + 100
	p.stock = 10
}
