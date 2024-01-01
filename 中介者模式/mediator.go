package main

type mediator interface {
	Set(Colleague1 colleague, Colleague2 colleague)
	Send(message string, Colleague colleague)
}

type colleague interface {
	setMediator(m mediator)
	Send(message string)
	Notify(message string)
}

type ConcreteMediator struct {
	c1 colleague
	c2 colleague
}

func (c *ConcreteMediator) Set(Colleague1 colleague, Colleague2 colleague) {
	c.c1 = Colleague1
	c.c2 = Colleague2
}

func (c *ConcreteMediator) Send(message string, Colleague colleague) {
	if c.c1 == Colleague {
		c.c2.Notify(message)
	} else {
		c.c1.Notify(message)
	}
}

type ConcreteColleagueA struct {
	m mediator
}

func (ca *ConcreteColleagueA) Notify(message string) {
	println("ca 收到信息：" + message)
}

func (ca *ConcreteColleagueA) Send(message string) {
	ca.m.Send(message, ca)
}

func (ca *ConcreteColleagueA) setMediator(m mediator) {
	ca.m = m
}

type ConcreteColleagueB struct {
	m mediator
}

func (cb *ConcreteColleagueB) Send(message string) {
	cb.m.Send(message, cb)
}

func (cb *ConcreteColleagueB) setMediator(m mediator) {
	cb.m = m
}

func (cb *ConcreteColleagueB) Notify(message string) {
	println("cb 收到信息：" + message)
}

func main() {
	ca := ConcreteColleagueA{}
	cb := ConcreteColleagueB{}

	m := ConcreteMediator{}
	m.Set(&ca, &cb)
	ca.setMediator(&m)
	cb.setMediator(&m)

	ca.Send("12345")
	cb.Send("67890")
}
