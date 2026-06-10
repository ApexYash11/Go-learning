package main

import "fmt"

func main() {
	 arr := []int{1, 2, 3, 4, 5}
	 sl1 := arr[0:3] // create a slice from index 0 to 2 (3 is exclusive)
	 sl2 := arr[2:5] // create a slice from index 2 to 4 (5 is exclusive)
	 fmt.Println("Slice 1:", sl1)
	 fmt.Println("Slice 2:", sl2)
	 sl  := arr[1:4]// create a slice from index 1 to 3 (4 is exclusive) 1 here is pointer to the original array and 4 is the length of the slice
	 sl3  :=sl[:2] // create a slice from index 0 to 1 (2 is exclusive) 0 here is pointer to the original array and 2 is the length of the slice
	 sl[0] = 10
	 fmt.Println("Array:", arr) // print the original array
	 fmt.Println("Slice:", sl,len(sl),cap(sl)) // print the slice, its length and capacity
	 fmt.Println("Slice 3:", sl3, len(sl3), cap(sl3)) // print the third slice, its length and capacity
}