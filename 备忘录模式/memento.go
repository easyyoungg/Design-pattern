package main

type Originator struct {
	state string
}

func (o *Originator) Get() string {
	return o.state
}

func (o *Originator) Storage() Memento {
	c := &Caretaker{}
	c.store(o.state)
	return c
}

func (o *Originator) Restore(m Memento) {
	o.state = m.read()
}

type Memento interface {
	store(state string)
	read() string
}

type Caretaker struct {
	state string
}

func (c *Caretaker) store(state string) {
	c.state = state
}

func (c *Caretaker) read() string {
	return c.state
}

func main() {
	o := Originator{state: "start"}
	println(o.Get())

	c := o.Storage()
	o.state = "end"
	println(o.Get())

	o.Restore(c)
	println(o.Get())
}
