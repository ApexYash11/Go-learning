package main

import "fmt"

func add[T int | float64 | uint](a, b T) T {
	return a + b
}

// genric are used to create functions that can work with different types of data. In this example, the add function can work with int, float64, and uint types. The getValue function is a generic function that takes a map with keys of type k and values of type V, and returns a slice of values of type V.

func getValue[k comparable , V any] (mp map [k]V) []V{
	values := []V{}
	for _, v := range mp {
		values = append(values, v)
	}
	return values
}

func main() {
  value := add(5, 10) // This will work with integers
	value2 := add(3.5, 4.5) // This will work with float64
	value3 := add(uint(5), uint(10)) // This will work with uint
	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	values := getValue(mp)
	fmt.Println("Values from map:", values)
	fmt.Println("Sum of unsigned integers:", value3)
	fmt.Println("Sum of floats:", value2)
	fmt.Println("Sum of integers:", value)	
}