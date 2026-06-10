package main

import "fmt"

func main() {
	x := 42
	fmt.Println("Hello", 2, x) // prints "Hello 2 42" with spaces in between
	fmt.Printf("%T",x)// formatted string output
	fmt.Printf("the value of x is %v :",x)
	y:=fmt.Sprintf("\"%10 .2f%%",x)
	fmt.Println(y)	
}