package builder

import "fmt"

/*
Use builder pattern to design parts of car using a fluent interface.
Builder type should return Car on calling build. Other methods of builer has to be fluent (return pointer to itself)
Can use abstract builder and different concrete builder if we have multiple implementation of builder
*/

type Car struct {
	tyre  int
	seat  int
	model string
}

type Builder struct {
	c Car
}

func NewBuilder() *Builder {
	return &Builder{c: Car{}}
}

func (b *Builder) SetTyre(t int) *Builder {
	b.c.tyre = t
	return b
}

func (b *Builder) SetSeat(s int) *Builder {
	b.c.seat = s
	return b
}

func (b *Builder) SetModel(m string) *Builder {
	b.c.model = m
	return b
}

func (b *Builder) Build() *Car {
	return &b.c
}

func Client() {
	b := NewBuilder()
	c := b.SetModel("sedan").SetSeat(2).SetTyre(4).Build()
	fmt.Println(c)
}

func Run() {
	Client()
}
