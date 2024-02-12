package main

import (
	"behavioral/internal/chainofresponsibility"
	"fmt"
)

func main() {
	fmt.Println("Chain of responsibility")
	chainofresponsibility.Run()
	fmt.Println("Done")
}
