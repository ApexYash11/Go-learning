package main

import (
	"fmt"
	"errors"
)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}


func deferFunc() { // This function will be deferred and executed when the surrounding function returns
	fmt.Println("defer")
	r := recover() // recover is used to catch the panic and prevent the program from crashing
	if r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func main() {
	defer deferFunc() // This will be executed even if a panic occurs
	// defer is used to ensure that the function is executed at the end of the surrounding function, even if a panic occurs
	panic("this caused an error") // This will cause a panic and the program will terminate after executing the deferred function
	// panic is runtime error, it will stop the execution of the program and print the error message
	fmt.Println("run")

	result,err := divide(10, 0) // This will cause an error because of division by zero
	if err != nil {
		fmt.Println("Error:", err) // This will print the error message
	} else {
		fmt.Println("Result:", result) // This will print the result if there is no error
	}
}