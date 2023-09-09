package main

import "fmt"

type FacadeInterface interface {
	MethodOne()
	MethodTwo()
}

type Facade struct{}

func (f Facade) MethodOne() {
	ca := ConcreteA{}
	ca.MethodA()
	cb := ConcreteB{}
	cb.MethodB()
}

func (f Facade) MethodTwo() {
	ca := ConcreteA{}
	ca.MethodA()
	cb := ConcreteC{}
	cb.MethodC()
}

type ConcreteA struct{}

func (ca ConcreteA) MethodA() {
	fmt.Println("it's ConcreteA")
}

type ConcreteB struct{}

func (cb ConcreteB) MethodB() {
	fmt.Println("it's ConcreteB")
}

type ConcreteC struct{}

func (cc ConcreteC) MethodC() {
	fmt.Println("it's ConcreteC")
}

func main() {
	f := Facade{}
	f.MethodTwo()
	f.MethodOne()
}
