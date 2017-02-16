package main

import (
	"fmt"
	"github.com/david7482/go-plugin-playground/calc"
)

func main()  {
	fmt.Print("This is a Go Application.\n")
	calc.SayHello("World")
	fmt.Printf("Add(3, 5): %d\n", calc.Add(3, 5))
}
