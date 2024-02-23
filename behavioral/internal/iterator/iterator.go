package iterator

import "fmt"

/* Parse elements sequentialy in an aggregated object

Collection -> interface
	-> createIterator
Iterator-> interface
	-> holds collection and index
	-> hasnext, next

Concreate for collection
	create concreate iterator

ConcreteIterator
	Has concreate collection
	index=0

hasnext => check index and len
next => return curr, index++

*/

type Collection interface {
	CreateIterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type IntCollection struct {
	c []int
}

func (i *IntCollection) CreateIterator() Iterator {
	return &NumberIterator{
		c:     i,
		index: 0,
	}
}

type NumberIterator struct {
	c     *IntCollection
	index int
}

func (i *NumberIterator) HasNext() bool {
	return i.index < len(i.c.c)
}

func (i *NumberIterator) Next() interface{} {
	val := i.c.c[i.index]
	i.index++
	return val
}

func Run() {
	Client()
}

func Client() {
	iC := IntCollection{c: []int{1, 2, 3}}
	it := iC.CreateIterator()
	for it.HasNext() {
		fmt.Println(it.Next().(int))
	}
}
