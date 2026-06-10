package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}	else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is less than 5")
	}

	a := 3
	switch  {
	case 1 <= a && a <= 1 :
		fmt.Println("a is 1")
	case 2 <= a && a <= 2 :
		fmt.Println("a is 2")
	default :
		fmt.Println("a is not 1 or 2")
	}
}