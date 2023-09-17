package main

import "fmt"

type director interface {
	build(b builder)
}

type directors struct{}

func (d directors) build(b builder) {
	b.StepOne()
	b.StepTwo()
	b.StepLast()
}

type builder interface {
	StepOne()
	StepTwo()
	StepLast()
}

type concreteBuilderA struct {
	isStepOneFinish bool
	isStepTwoFinish bool
}

func (ca *concreteBuilderA) StepOne() {
	ca.isStepOneFinish = true
	fmt.Println("concreteBuilderA step one is finish.")
}

func (ca *concreteBuilderA) StepTwo() {
	ca.isStepTwoFinish = true
	fmt.Println("concreteBuilderA step two is finish.")
}

func (ca *concreteBuilderA) StepLast() {
	if ca.isStepTwoFinish && ca.isStepOneFinish {
		fmt.Println("concreteBuilderA is finish!")
	} else {
		fmt.Println("concreteBuilderA build unsuccess.")
	}
}

type concreteBuilderB struct {
	isStepOneFinish bool
	isStepTwoFinish bool
}

func (cb *concreteBuilderB) StepOne() {
	cb.isStepOneFinish = true
	fmt.Println("concreteBuilderB step one is finish.")
}

func (cb *concreteBuilderB) StepTwo() {
	cb.isStepTwoFinish = true
	fmt.Println("concreteBuilderB step two is finish.")
}

func (cb *concreteBuilderB) StepLast() {
	if cb.isStepTwoFinish && cb.isStepOneFinish {
		fmt.Println("concreteBuilderB is finish!")
	} else {
		fmt.Println("concreteBuilderB build unsuccess.")
	}
}

func main() {
	b := directors{}
	ca := concreteBuilderA{}
	b.build(&ca)
	cb := concreteBuilderB{}
	b.build(&cb)
}
