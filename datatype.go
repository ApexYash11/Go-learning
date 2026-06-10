package main

import "fmt"

func main() {
	int:= 42
	var x	string = "Hello, Go!"
	const pi float64 = 3.14
	y := 3.14
	fmt.Println("%T", y)
	fmt.Println("The value of pi is :", pi)
	fmt.Println("The value of x is:", x)
	fmt.Println("The value of int is:", int)
}