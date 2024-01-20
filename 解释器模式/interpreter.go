package main

import "strings"

type AbstractExprssion interface {
	Interpret(s string)
}

type TerminalExprssion struct{}

type NonTerminalExprssion struct{}

func (t TerminalExprssion) Interpret(s string) {
	if strings.Contains(s, "Y") {
		println("TerminalExprssion.")
	}
}

func (n NonTerminalExprssion) Interpret(s string) {
	if strings.Contains(s, "N") {
		println("NonTerminalExprssion.")
	}
}

func main() {
	str1 := "N"
	str2 := "Y"
	t := TerminalExprssion{}
	n := NonTerminalExprssion{}
	t.Interpret(str1)
	t.Interpret(str2)

	n.Interpret(str1)
	n.Interpret(str2)
}
