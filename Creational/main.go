package main

import (
	"creational/abstractFactory"
	"creational/factory"
	"fmt"
)

func main() {
	fmt.Println("Factory")
	factory.Run()
	fmt.Println("Abstract Factory")
	abstractFactory.Run()
	fmt.Println("Done")
}
