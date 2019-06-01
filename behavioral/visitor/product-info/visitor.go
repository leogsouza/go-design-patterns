package main

import (
	"fmt"
)

// ProductInfoRetriever contains methods to return info
// about a product
type ProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

// Visitor gets product information
type Visitor interface {
	Visit(ProductInfoRetriever)
}

// Visitable accepts a Visitor
type Visitable interface {
	Accept(Visitor)
}

// Product represents a product
// This struct has information about Price and Name
type Product struct {
	Price float32
	Name  string
}

// GetPrice returns the product's price
func (p *Product) GetPrice() float32 {
	return p.Price
}

// GetName returns the product's name
func (p *Product) GetName() string {
	return p.Name
}

// Accept puts the product in Visitor list
func (p *Product) Accept(v Visitor) {
	v.Visit(p)
}

// Rice is a type of product
type Rice struct {
	Product
}

// Pasta is a type of product
type Pasta struct {
	Product
}

// Fridge is a type of product
type Fridge struct {
	Product
}

// GetPrice returns the fridge product's price
func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

// Accept puts the fridge product into visitor list
func (f *Fridge) Accept(v Visitor) {
	v.Visit(f)
}

// PriceVisitor holds the sum of products
type PriceVisitor struct {
	Sum float32
}

// Visit sum the price of all products
func (pv *PriceVisitor) Visit(p ProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

// NamePrinter is a struct to print all products
type NamePrinter struct {
	ProductList string
}

// Visit prints the information of all products
func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

func main() {
	products := make([]Visitable, 3)
	products[0] = &Rice{
		Product: Product{
			Price: 32.0,
			Name:  "Some rice",
		},
	}

	products[1] = &Pasta{
		Product: Product{
			Price: 40.0,
			Name:  "Some pasta",
		},
	}

	products[2] = &Fridge{
		Product: Product{
			Price: 50,
			Name:  "A fridge",
		},
	}

	// Print the sum of prices
	priceVisitor := &PriceVisitor{}

	for _, p := range products {
		p.Accept(priceVisitor)
	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)

	// Print the products list
	nameVisitor := &NamePrinter{}

	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n-----------\n%s",
		nameVisitor.ProductList)
}
