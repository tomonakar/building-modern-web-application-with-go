package main

import (
	"log"
	"sort"
)

func main() {
	// Sliceの宣言
	var mySlice []string
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")
	log.Println(mySlice)

	// Sliceの宣言
	var mySlice2 []int
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)
	log.Println(mySlice2)
	sort.Ints(mySlice2)
	log.Println(mySlice2)

	// Slice宣言のショートハンド
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[0:2])

	names := []string{"one", "seven", "fish", "red", "blue", "two"}
	log.Println(names)
}
