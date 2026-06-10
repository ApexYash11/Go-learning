package main 

import "fmt"

func main() {
   var arr [5]int // array of 5 integers

	 arr1 := [3]int{1, 2, 3} // array of 3 integers with values
	 arr2 := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	 arr3 := [...][1]int {{1}, {4}, {7}} // array of arrays with inferred size0
	 arr1[0] = 10 // set the first element of arr1 to 10
	 arr2[0][0] = 20 // set the first element of the first sub-array of arr2 to 20
	 arr3[0][0] = 30 // set the first element of the first sub-array of arr3 to 30

	 for i, v := range arr1 { // iterate over arr1 using range
	 	fmt.Printf("Index: %d, Value: %d\n", i, v) // print the index and value of each element in arr1
	 }

	 fmt.Println(arr) // print the array
	 fmt.Println(arr1) // print the second array
	 fmt.Println(arr2) // print the third array
	 fmt.Println(arr3) // print the fourth array
	 fmt.Println("Length of arr1:", len(arr1)) // print the length of arr1
	 fmt.Println("Length of arr2:", len(arr2))

}