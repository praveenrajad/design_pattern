package factory

/*
interface: AbsProduct
	struct: ConceteProductA
	struct: ConcreteProductB

func Factory: return AbsProduct

Client:
	cPA := ConceteProductA{}
	pritnDetail(cpA)
func printDetail (AbsProdutc)

Learning: Interface is already pass by reference, so not need to use pointers

*/

import "fmt"

func Run() {
	Client()
}

type AbsProduct interface {
	Print()
}

type ConcreteProductA struct {
	name string
}

func (pA *ConcreteProductA) Print() {
	fmt.Println(pA.name)
}

type ConcreteProductB struct {
	name string
}

func (pB *ConcreteProductB) Print() {
	fmt.Println(pB.name)
}

func ProductFactory(name string) AbsProduct {
	if name == "A" {
		return &ConcreteProductA{"A"}
	} else if name == "B" {
		return &ConcreteProductB{"B"}
	}
	return nil
}

func Client() {
	productA := ProductFactory("A")
	printDetails(productA)
	productB := ProductFactory("B")
	printDetails(productB)
}

func printDetails(p AbsProduct) {
	p.Print()
}
