package main

import "fmt"

type subject interface {
	sell()
	rent()
}

type Proxy struct{}

func (p Proxy) sell() {
	var r realsubject
	r.sell()
}

func (p Proxy) rent() {
	var r realsubject
	r.rent()
}

type realsubject struct{}

func (r realsubject) sell() {
	fmt.Printf("is sell.\n")
}

func (r realsubject) rent() {
	fmt.Printf("is rent.\n")
}

func main() {
	var p Proxy

	p.rent()
	p.sell()
}
