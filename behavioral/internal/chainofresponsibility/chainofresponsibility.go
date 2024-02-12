package chainofresponsibility

import "fmt"

func Run() {
	Client()
}

type Handler interface {
	SetNext(Handler)
	Handler(amount int)
}

type Manager struct {
	nextHandler Handler
}

func (m *Manager) SetNext(h Handler) {
	m.nextHandler = h
}

func (m *Manager) Handler(amount int) {
	if amount > 5000 {
		fmt.Println("Calling next from manager")
		m.nextHandler.Handler(amount)
		return
	}
	fmt.Println("Manager handled it")
}

type Director struct {
	nextHandler Handler
}

func (d *Director) SetNext(h Handler) {
	d.nextHandler = h
}

func (d *Director) Handler(amount int) {
	if amount > 10000 {
		fmt.Println("Unhandled")
		return
	}
	fmt.Println("Director handled it")
}

func Client() {
	m := &Manager{}
	d := &Director{}
	m.SetNext(d)
	m.Handler(1000)
	fmt.Println("------")
	m.Handler(6000)
	fmt.Println("-----")
	m.Handler(11000)
}
