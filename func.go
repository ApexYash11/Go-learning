package main

import "fmt"

func add(a, b int) (int, string) {
	return a + b, "success"
}

func sums(nums...int) int{
	s := 0
	for _, num := range nums {
		s += num
	}
	return s

}

func funct(callable func(int)int ) int{
	return callable(10)
}

func getfunc(str string) func(string) string{
	return func(str2 string) string{
		return str + " " + str2
	}
}

func main() {
	value, message := add(3, 5) // call the add function with arguments 3 and 5
	value1 := funct(func(x int ) int{
		return x+1
	})

	f1 :=getfunc("Hello")
	s := sums(1,2,3,4,5)
	fmt.Println("The sum of the numbers is:", s) // print the sum of the numbers
	value2 := f1("World")
	value3 := getfunc("Goodbye")("Everyone")
	fmt.Println(value2)
	fmt.Println(value3)
	fmt.Println(value1)
	fmt.Println("The sum is:", value) // print the result
	fmt.Println("Message:", message) // print the message
}