package main

import "fmt"

type Aggregate interface {
	CreateInterator() Iterator
}

type Iterator interface {
	Set(a []string)
	First() string
	Next() string
	IsDone() bool
	CurrentItem() string
}

type ConcreteIterator struct {
	count int
	list  []string
}

func (ci *ConcreteIterator) Set(a []string) {
	ci.list = a
}

func (ci *ConcreteIterator) First() string {
	ci.count = 0
	return ci.list[0]
}

func (ci *ConcreteIterator) Next() string {
	ci.count++
	if ci.IsDone() {
		return ""
	}
	return ci.list[ci.count]
}

func (ci *ConcreteIterator) IsDone() bool {
	return ci.count >= len(ci.list)
}

func (ci *ConcreteIterator) CurrentItem() string {
	return ci.list[ci.count]
}

type ConcreteAggregate struct{}

func (ca ConcreteAggregate) CreateInterator() Iterator {
	return &ConcreteIterator{}
}

func main() {
	ca := ConcreteAggregate{}
	ci := ca.CreateInterator()
	ci.Set([]string{"1", "2", "3"})
	ci.First()
	for !ci.IsDone() {
		fmt.Printf("%s is in iterator\n", ci.CurrentItem())
		ci.Next()
	}
}
