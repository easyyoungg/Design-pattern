package main

import "fmt"

type State interface {
	Do()
}

type Context struct {
	State
}

func (c *Context)Set(s State)  {
	c.State = s
}

func (c *Context)Get() State {
	return c.State
}

type StateA struct {}

func (s StateA)Do()  {
	fmt.Println("stateA is doing.")
}

type StateB struct {}

func (s StateB)Do()  {
	fmt.Println("stateB is doing.")
}

func main()  {
	var c Context
	var sa StateA
	c.Set(sa)
	c.Get().Do()
	var sb StateB
	c.Set(sb)
	c.Get().Do()
}

