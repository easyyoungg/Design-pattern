package main

import "fmt"

type Abstrct interface {
	Template()
	Method()
}

type AbstrctClass struct{}

func (a AbstrctClass) Template() {
	fmt.Println("Template.")
}

type ConcreteClassA struct{}

func (ca ConcreteClassA) Template() {
	a := AbstrctClass{}
	a.Template()
}

func (ca ConcreteClassA) Method() {
	fmt.Println("ConcreteClassA.")
}

type ConcreteClassB struct{}

func (cb ConcreteClassB) Template() {
	a := AbstrctClass{}
	a.Template()
}

func (cb ConcreteClassB) Method() {
	fmt.Println("ConcreteClassB.")
}

func main() {
	ca := ConcreteClassA{}
	ca.Template()
	ca.Method()
	cb := ConcreteClassB{}
	cb.Template()
	cb.Method()
}
