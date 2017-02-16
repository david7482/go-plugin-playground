package calc

import "fmt"

func SayHello(name string) {
	fmt.Printf("Go says: Hello, %s!\n", name)
}

func Add(num0, num1 int) int {
	return num0 + num1
}