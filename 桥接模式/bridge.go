package main

type Abstraction interface {
	Operation()
	Implementor
}

type RefinedAbstraction struct {
	Implementor
}

func (r RefinedAbstraction) Operation() {
	println("diff part.")
}

type Implementor interface {
	OperationImp()
}

type ConcreteImplementorA struct{}

type ConcreteImplementorB struct{}

func (ca ConcreteImplementorA) OperationImp() {
	println("Common partA.")
}

func (cb ConcreteImplementorB) OperationImp() {
	println("Common partB.")
}

func main() {
	ra := RefinedAbstraction{}

	ra.Implementor = ConcreteImplementorA{}
	ra.Operation()
	ra.OperationImp()

	ra.Implementor = ConcreteImplementorB{}
	ra.Operation()
	ra.OperationImp()
}
