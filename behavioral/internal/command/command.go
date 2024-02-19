package command

import "fmt"

/*
Seperate the command invoker and the executer are seperate. So execution can be modified
invoker -> call (set and execute command)
	command  => inteface
receiver => actual command gets executed

Invoker -> set and execute command
command interface
	concrete commandA & commandB
Receiver=> called from concreateCommand where actual command gets executed.

Having a seperate receiver or executing witin the command itself is design based
*/

func Run() {
	Client()
}

type Receiver struct {
}

func (r *Receiver) Execute() {
	fmt.Println("command got executed")
}

type Command interface {
	Execute()
}

type CommandA struct {
	r *Receiver
}

func (c *CommandA) Execute() {
	fmt.Println("Executing command A")
	c.r.Execute()
}

type Invoker struct {
	c Command
}

func (i *Invoker) SetCommand(c Command) {
	i.c = c
}

func (i *Invoker) Execute() {
	i.c.Execute()
}

func Client() {
	cA := &CommandA{&Receiver{}}
	i := Invoker{}
	i.SetCommand(cA)
	i.Execute()
}
