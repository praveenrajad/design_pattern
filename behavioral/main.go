package main

import (
	"behavioral/internal/chainofresponsibility"
	"behavioral/internal/command"
	"fmt"
)

func main() {
	fmt.Println("Chain of responsibility")
	chainofresponsibility.Run()
	fmt.Println("Done")
	fmt.Println("Command")
	command.Run()
	fmt.Println("Done")
}
