package main

import "fmt"

const StrategyAName = "StrategyA"
const StrategyBName = "StrategyB"

type Context struct {
	s Strategy
}

func (c Context) Set(str string) {
	var s Strategy
	switch str {
	case StrategyAName:
		s = StrategyA{}
	case StrategyBName:
		s = StrategyB{}
	}
	c.s = s
	c.s.Do()
}

func (c Context) SetNoFactory(s Strategy) {
	c.s = s
	c.s.Do()
}

type Strategy interface {
	Do()
}

type StrategyA struct{}

type StrategyB struct{}

func (s StrategyA) Do() {
	fmt.Printf("StrategyA do.\n")
}

func (s StrategyB) Do() {
	fmt.Printf("StrategyB do.\n")
}

func main() {
	var c Context
	var sa StrategyA
	var sb StrategyB
	c.Set(StrategyAName)
	c.Set(StrategyBName)
	c.SetNoFactory(sa)
	c.SetNoFactory(sb)
}
