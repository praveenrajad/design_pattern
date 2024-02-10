package singleton

import (
	"fmt"
	"sync"
)

func Run() {
	Client()
}

type single struct {
}

var (
	instance *single
	once     sync.Once
)

func getSingle() *single {
	once.Do(func() {
		instance = &single{}
	})
	return instance
}

func getInstance(wg *sync.WaitGroup) {
	defer wg.Done()
	i := getSingle()
	fmt.Printf("address of %p \n", i)

}

func Client() {
	var wg sync.WaitGroup
	wg.Add(2)
	go getInstance(&wg)
	go getInstance(&wg)
	wg.Wait()

}
