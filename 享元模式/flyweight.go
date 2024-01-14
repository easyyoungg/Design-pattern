package main

import "fmt"

type Flyweight interface {
	Operation()
}

type ConcreteFlyweight struct {
	name string
}

func (c *ConcreteFlyweight) Operation() {
	fmt.Println("share flyweight:" + c.name)
}

type UnsharedConcreteFlyweight struct {
	name string
}

func (c *UnsharedConcreteFlyweight) Operation() {
	fmt.Println("Unshare flyweight:" + c.name)
}

type FlyweightFactory struct {
	m map[string]struct{}
}

func (f *FlyweightFactory) GetFlyweight(key string) ConcreteFlyweight {
	if _, ok := f.m[key]; !ok {
		f.m[key] = struct{}{}
	}
	return ConcreteFlyweight{name: key}
}

func main() {
	f := FlyweightFactory{m: make(map[string]struct{})}

	x := f.GetFlyweight("x")
	x.Operation()

	y := f.GetFlyweight("y")
	y.Operation()

	uf := UnsharedConcreteFlyweight{name: "z"}
	uf.Operation()

	newx := f.GetFlyweight("x")
	newx.Operation()
}
