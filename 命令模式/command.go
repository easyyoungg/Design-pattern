package main

type Command interface {
	Excute()
}

type ConcreteCommandA struct {
	Receiver
}

func (c ConcreteCommandA) Excute() {
	c.Receiver.ExcuteCommandA()
}

type ConcreteCommandB struct {
	Receiver
}

func (c ConcreteCommandB) Excute() {
	c.Receiver.ExcuteCommandB()
}

type Receiver struct{}

func (r Receiver) ExcuteCommandA() {
	println("Excute commandA.")
}

func (r Receiver) ExcuteCommandB() {
	println("Excute commandB.")
}

type Invoker struct {
	c Command
}

func (i *Invoker) Set(c Command) {
	i.c = c
}

func (i *Invoker) Do() {
	i.c.Excute()
}

func main() {
	i := Invoker{}
	ca := ConcreteCommandA{}
	cb := ConcreteCommandB{}

	i.Set(ca)
	i.Do()

	i.Set(cb)
	i.Do()
}
