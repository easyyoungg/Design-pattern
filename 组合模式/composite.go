package main

import "fmt"

type Component interface {
	Add(cp Component)
	Delete(cp Component)
	Name() string
	Display()
}

type Composite struct {
	cs   []Component
	name string
}

func (c *Composite) Add(cp Component) {
	c.cs = append(c.cs, cp)
}

func (c *Composite) Delete(cp Component) {
	for i, com := range c.cs {
		if com.Name() == cp.Name() {
			c.cs = append(c.cs[:i], c.cs[i+1:]...)
		}
	}
}

func (c *Composite) Display() {
	result := ""
	for _, com := range c.cs {
		result += com.Name()
	}
	fmt.Printf("the composite it's contain %s\n", result)
}

func (c *Composite) Name() string {
	return c.name
}

type Leaf struct {
	name string
}

func (l *Leaf) Add(cp Component) {} //no mean

func (l *Leaf) Delete(cp Component) {} //no mean

func (l *Leaf) Display() {
	fmt.Printf("it's leaf %s\n", l.name)
}

func (l *Leaf) Name() string {
	return l.name
}

func main() {
	l1 := &Leaf{name: "l1"}
	l2 := &Leaf{name: "l2"}
	c1 := &Composite{name: "c1"}
	c2 := &Composite{name: "c2"}
	c1.Add(l1)
	c1.Add(c2)
	c2.Add(l2)
	l1.Display()
	l2.Display()
	c1.Display()
	c2.Display()
}
