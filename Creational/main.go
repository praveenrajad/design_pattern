package main

import (
	"creational/internal/abstractFactory"
	"creational/internal/builder"
	"creational/internal/clone"
	"creational/internal/factory"
	"fmt"
)

func main() {
	fmt.Println("Factory")
	factory.Run()
	fmt.Println("Abstract Factory")
	abstractFactory.Run()
	fmt.Println("Done")
	fmt.Println("Builder")
	builder.Run()
	fmt.Println("Done")
	fmt.Println("Clone")
	clone.Run()
	fmt.Println("Done")
}
