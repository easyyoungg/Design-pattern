package main

import "fmt"

const T1NAME = "T1"
const T2NAME = "T2"

const A1NAME = "A1"
const A2NAME = "A2"

type Product interface {
	Do()
}

type T1 struct {}
type T2 struct {}

func (t T1)Do()  {
	fmt.Printf("T1 do.\n")
}

func (t T2)Do()  {
	fmt.Printf("T2 do.\n")
}

type A1 struct {}
type A2 struct {}

func (a A1)Do()  {
	fmt.Printf("A1 do.\n")
}

func (a A2)Do()  {
	fmt.Printf("A2 do.\n")
}

type AbstractFactory interface {
	Get(s string) Product
}

type AFactory struct {}
type TFactory struct {}

func (a AFactory)Get(s string) Product {
	switch s {
	case A1NAME:
		return A1{}
	case A2NAME:
		return A2{}
	}
	return nil
}

func (a TFactory)Get(s string) Product {
	switch s {
	case T1NAME:
		return T1{}
	case T2NAME:
		return T2{}
	}
	return nil
}

func main()  {
	var t TFactory
	var a AFactory
	t.Get(T1NAME).Do()
	t.Get(T2NAME).Do()
	a.Get(A1NAME).Do()
	a.Get(A2NAME).Do()
}
