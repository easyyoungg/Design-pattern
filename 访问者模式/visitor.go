package main

type Element interface {
	Accept(v visitor)
}

type ConcreteElementA struct{}

type ConcreteElementB struct{}

func (ca *ConcreteElementA) Accept(v visitor) {
	v.VisitConcreteElementA()
}

func (cb *ConcreteElementB) Accept(v visitor) {
	v.VisitConcreteElementB()
}

type visitor interface {
	VisitConcreteElementA()
	VisitConcreteElementB()
}

type ConcreteVisitorA struct{}

type ConcreteVisitorB struct{}

func (ca ConcreteVisitorA) VisitConcreteElementA() {
	println("visitorA on ElementA.")
}

func (ca ConcreteVisitorA) VisitConcreteElementB() {
	println("visitorA on ElementB.")
}

func (cb ConcreteVisitorB) VisitConcreteElementA() {
	println("visitorB on ElementA.")
}

func (cb ConcreteVisitorB) VisitConcreteElementB() {
	println("visitorB on ElementB.")
}

type ObjectStructure struct{}

func (o ObjectStructure) Display(v visitor) {
	ca := ConcreteElementA{}
	cb := ConcreteElementB{}

	ca.Accept(v)
	cb.Accept(v)
}

func main() {
	o := ObjectStructure{}
	va := ConcreteVisitorA{}
	vb := ConcreteVisitorB{}
	o.Display(va)
	o.Display(vb)
}
