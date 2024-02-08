package compositionforsamefunc

import "fmt"

/*
Learn: If we need to reuse a func across different structs, we can embed another type in these concreate impementation
Note: if Print function is redefined in product C then the method is overriden, and effectively has the overriden Print
*/

func Run() {
	Client()
}

type Printer interface {
	Print()
}

type BasePrinter struct{}

func (b *BasePrinter) Print() {
	fmt.Println("Hello")
}

type AbsProduct interface {
	Printer
}

type ProductA struct {
	BasePrinter
}

type ProductB struct {
	BasePrinter
}

type ProductC struct {
	BasePrinter
}

func (p *ProductC) Print() {
	fmt.Println("Printing from C")
}

func Client() {
	var p AbsProduct
	p = &ProductA{}
	p.Print()
	p = &ProductB{}
	p.Print()
	p = &ProductC{}
	p.Print()
}
