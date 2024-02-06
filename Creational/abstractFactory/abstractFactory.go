package abstractFactory

import "fmt"

/*
Here the factory is abstracted. 
AbsFactory create multiple related AbstractProducts.
For implementation we need to have ConcreateFactories and Concrete products
in client we create concrete factories and then individual concrete products

Learning: Interace method signature should exactly match (we cant use args/return with concreate struct instead of interface)
*/

func Run() {
	Client()
}

type AbsProductA interface {
	Print()
}
type AbsProductB interface {
	Print()
}

type ConcreteProductACompany1 struct {
	name string
}

func (p *ConcreteProductACompany1) Print() {
	fmt.Println(p.name)
}

type ConcreteProductBCompany1 struct {
	name string
}

func (p *ConcreteProductBCompany1) Print() {
	fmt.Println(p.name)
}

type ConcreteProductACompany2 struct {
	name string
}

func (p *ConcreteProductACompany2) Print() {
	fmt.Println(p.name)
}

type ConcreteProductBCompany2 struct {
	name string
}

func (p *ConcreteProductBCompany2) Print() {
	fmt.Println(p.name)
}

type AbsFactoryCompany interface {
	CreateProductA() AbsProductA
	CreateProductB() AbsProductB
}

type FactoryCompany1 struct {
	name string
}

func (f1 *FactoryCompany1) CreateProductA() AbsProductA {
	return &ConcreteProductACompany1{"Company 1, Product A"}
}

func (f1 *FactoryCompany1) CreateProductB() AbsProductB {
	return &ConcreteProductBCompany1{"Company 1, Product B"}
}

type FactoryCompany2 struct {
	name string
}

func (f2 *FactoryCompany2) CreateProductA() AbsProductA {
	return &ConcreteProductACompany2{"Company 2, Product A"}
}

func (f2 *FactoryCompany2) CreateProductB() AbsProductB {
	return &ConcreteProductBCompany2{"Company2, Product B"}
}

func Client() {
	f1 := FactoryCompany1{"Factory 1"}
	f2 := FactoryCompany2{"Factory 2"}
	fmt.Println(f1, f2)
	f1pA := f1.CreateProductA()
	f1pB := f1.CreateProductB()
	f2pA := f2.CreateProductA()
	f2pB := f2.CreateProductA()
	printDetails(f1pA, f1pB)
	printDetails(f2pA, f2pB)
}


func printDetails(a AbsProductA, b AbsProductB) {
	a.Print()
	b.Print()
}