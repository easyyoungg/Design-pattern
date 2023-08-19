package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
	age  int
}

func (c ConcretePrototype) Clone() Prototype {
	d := c
	return d
}

func main() {
	var c = &ConcretePrototype{name: "xiaoming", age: 18}
	d := c.Clone()
	ds, ok := d.(ConcretePrototype)
	fmt.Println(ds.age, ds.name, ok)
}
