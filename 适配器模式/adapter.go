package main

import "fmt"

type target interface {
	request()
}

type adapter struct{}

func (a adapter) request() {
	var adp adaptee
	adp.otherrequest()
}

type adaptee struct{}

func (a adaptee) otherrequest() {
	fmt.Println("This is otherrequest.")
}

func main() {
	var a adapter
	a.request()
}
