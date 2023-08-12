package main

import "fmt"

const T1NAME = "T1"
const T2NAME = "T2"

type prodcut interface {
	Do()
}

type factory interface {
	Get() prodcut
}

type T1 struct{}
type T2 struct{}

func (t T1) Do() {
	fmt.Printf("T1 do.\n")
}

func (t T2) Do() {
	fmt.Printf("T2 do.\n")
}

type T1factory struct{}

func (t T1factory) Get() prodcut {
	return T1{}
}

type T2factory struct{}

func (t T2factory) Get() prodcut {
	return T2{}
}

func main() {
	var t1 T1factory
	var t2 T2factory
	t1.Get().Do()
	t2.Get().Do()
}
