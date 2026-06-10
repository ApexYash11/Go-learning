package main

import "fmt"

func main() {
	x := 10
	switch	{
	case x > 5 :
		fmt.Println("x is greater than 5")
		fallthrough // this will execute the next case even if the condition is not met, so it will print both "x is greater than 5" and "x is equal to 5"
	case x == 5 :
		fmt.Println("x is equal to 5")
		fallthrough
	default :
		fmt.Println("x is less than 5")
	}
}