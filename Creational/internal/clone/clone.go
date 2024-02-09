package clone

/*
There is interface Protype that has a method Clone. Any struct can impletment clone to return a deep copy of itself.
On the client the return type needs to be converted from interface to the concrete struct
*/

import "fmt"

func Run() {
	Client()
}

type Prototype interface {
	Clone() Prototype
}

type Bank struct {
	name    string
	count   int
	address string
}

func (b *Bank) Clone() Prototype {
	return &Bank{b.name, b.count, b.address}
}

func Client() {
	b1 := &Bank{"bank", 2, "chennai"}
	b2 := b1.Clone().(*Bank)
	fmt.Println("Bank1", b1)
	fmt.Println("Bank2", b2)
}
