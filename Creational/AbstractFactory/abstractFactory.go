package abstractFactory

/*

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
tye

type ConcreteProductACompany1 struct {
	name string
}

func (pac1 *ConcreteProductACompany1) Print()


type ConcreteProductBCompany1 struct {
	name string
}

type ConcreteProductACompany2 struct {
	name string
}

type ConcreteProductBCompany2 string {
	name string
}


type AbsFactoryCompany interface {	
	CreateProductA() AbsProductA
	CreateProductB() AbsProductB
}

type FactoryCompany1 struct {
	name string
}

func (f1 *FactoryCompany1) CreateProductA () ConcreteProductACompany1 {
	return ConcreteProductACompany1 {"Company 1, Product A"}
}

func (f1 *FactoryCompany1) CreateProductB() ConcreteProductBCompany1 {
	return ConcreteProductBCompany1{"Company 1, Product B"}
}


type FactoryCompany2 struct {
	name string
}

func(f2 *FactoryCompany2) CreateProductA () ConcreteProductACompany2 {
	return ConcreteProductACompany2{"Company 2, Product A"}
}

func (f2 *FactoryCompany2) CreateProductB() ConcreteProductACompany2 {
	return ConcreteProductBCompany2{"Company2, Product B"}
}

func Client() {
	f1 := FactoryCompany1{"Factory 1"}
	f2 := FactoryCompany2{"Factory 2"}
	f1pA := f1.CreateProductA()
	f1pB := f1.CreateProductB()
	f2pA := f2.CreateProductA()
	f2pB := f2.CreateProductA()
}