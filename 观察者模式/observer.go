package main

import "fmt"

type subject interface {
	Attach(o observer)
	Detach(o observer)
	Notify()
}

type observer interface {
	Update()
}

type concretesubjectA struct {
	os []observer
}

func (ca *concretesubjectA) Attach(o observer) {
	ca.os = append(ca.os, o)
}

func (ca *concretesubjectA) Detach(o observer) {
	for i, oo := range ca.os {
		if oo == o {
			ca.os = append(ca.os[:i], ca.os[i+1:]...)
		}
	}
}

func (ca *concretesubjectA) Notify() {
	for _, oo := range ca.os {
		oo.Update()
	}
}

type concretesubjectB struct {
	os []observer
}

func (cb *concretesubjectB) Attach(o observer) {
	cb.os = append(cb.os, o)
}

func (cb *concretesubjectB) Detach(o observer) {
	for i, oo := range cb.os {
		if oo == o {
			cb.os = append(cb.os[:i], cb.os[i+1:]...)
		}
	}
}

func (cb *concretesubjectB) Notify() {
	for _, oo := range cb.os {
		oo.Update()
	}
}

type concreteobserverA struct {
	name string
}

func (ca *concreteobserverA) Update() {
	fmt.Println(ca.name + " concreteobserverA is update.")
}

type concreteobserverB struct {
	name string
}

func (cb *concreteobserverB) Update() {
	fmt.Println(cb.name + " concreteobserverB is update.")
}

func main() {
	csa := &concretesubjectA{}
	csb := &concretesubjectB{}
	coa := &concreteobserverA{name: "xiaoming"}
	cob := &concreteobserverB{name: "xiaohong"}

	csa.Attach(coa)
	csa.Attach(cob)
	csa.Notify()

	csb.Attach(coa)
	csb.Notify()
}
