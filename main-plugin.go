package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("This is a Go Application");

	p, err := plugin.Open("calc-plugin.so")
	if err != nil {
		panic(err)
	}

	sayHelloSymbol, err := p.Lookup("SayHello")
	if err != nil {
		panic(err)
	}

	addSymbol, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}

	sayHelloSymbol.(func(string))("World")
	fmt.Printf("Add(3, 5): %d\n", addSymbol.(func(int, int)int)(3, 5))
}