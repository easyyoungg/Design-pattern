package main

import (
	"fmt"
)

var a = Singleton{}

type Singleton struct{}

func GetInstance() Singleton {
	return a
}

func main() {
	instance := GetInstance()
	fmt.Printf("%v", instance)
}
