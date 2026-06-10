package main

import ("fmt"
			"math"
			"strconv"
		)

func main() {
	x := 42
	y := 3.14
	c :="Hello, World!"
	b := "Go is awesome!"
	e := "1234"
	f,err := strconv.Atoi(e)
	z:= x + int(y) // type conversion from float64 to int
	// we do this beacuse go  confuse and does not allow to add different data types together, so better to convert smaller vlaue to bigger value, in this case we convert float64 to int
	a := float64(x) / float64(y) // type conversion from int to float64
	d :=c+b
	
  fmt.Println(f, err)
 	fmt.Println(math.Min(6,7))
	fmt.Println("The value of z is:", z)
	fmt.Println("The value of a is:", a)
	fmt.Println("The value of d is:", d)
}