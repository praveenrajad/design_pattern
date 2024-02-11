package objectpool

import "fmt"

/*

Learn: Chan can be used for a single thread. We create a chan of objects of a particuar size. Get empties chan and returns an object . Put fills the chan,
can be done until full.
*/

func Run() {
	Client()
}

type Object struct {
	id int
}

type Pool struct {
	c chan *Object
}

func (p *Pool) Get() *Object {
	select {
	case obj := <-p.c:
		return obj
	default:
		return nil
	}
}

func (p *Pool) Put(o *Object) {
	p.c <- o
}

func CreatePool(s int) *Pool {
	p := Pool{c: make(chan *Object, s)}
	for i := 0; i < s; i++ {
		obj := &Object{i}
		fmt.Println("Create: ", &obj)
		p.c <- obj
	}
	return &p
}

func Client() {
	p := CreatePool(2)
	obj1 := p.Get()
	fmt.Println("Got obj", obj1.id)
	obj2 := p.Get()
	fmt.Println("Got obj", obj2.id)
	obj3 := p.Get()
	fmt.Printf("Got obj %p\n", obj3)

	p.Put(obj1)
	p.Put(obj2)
	obj1 = p.Get()
	fmt.Println("Got obj", obj1.id)
	obj2 = p.Get()
	fmt.Println("Got obj", obj2.id)

}
