package main

import "C"
import "fmt"

func main() {}

//export SayHello
func SayHello(name string) {
	fmt.Printf("Go says: Hello, %s!\n", name)
}

//export Add
func Add(num0, num1 int) int {
	return num0 + num1
}