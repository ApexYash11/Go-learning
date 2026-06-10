package main

import "fmt"

func main() {
  sl := []string{"Go", "is", "awesome"}

	for i :=0 ; i<=10;i++{
		sl=append(sl,"yash")
		fmt.Println(sl,len(sl),cap(sl))
	}
	fmt.Println(sl) // print the slice
}