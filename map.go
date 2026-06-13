package main

import "fmt"

func main() {
	// mp := make(map[string]int) // create an empty map
	mp := map[string]int{"apple":  5, "banana": 3,"orange": 2,}	// create a map with string keys and int values using a map literal
	mp2 := map[string] []int{"grape": []int{4}, "melon": []int{6}} // create a map with string keys and slice of int values
	mp3 :=map[string] int {} // create an empty map using a map literal
	delete(mp3, "apple") // delete the key "apple" from mp3 (no effect since it doesn't exist)
	value ,ok := mp["banana"] // retrieve the value for the key "banana" and check if it exists
	fmt.Println(value, ok) // print the value and whether it exists
	fmt.Println("Map:", mp) // print the map
	fmt.Println("Map2:", mp2) // print the second map
	fmt.Println("Map3:", mp3) // print the third map

	mp4 := map[uint]uint{}
	n := uint(100)

	for i := uint(0); i < n; i++ {
		for d := uint(1); d <= 5; d++ {
	 if i%d == 0 {
				mp4[i]++
			}
		}

	}

	fmt.Println("Map4:", mp4) // print the fourth map

}
