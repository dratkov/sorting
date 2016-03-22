package main

import (
	"fmt"
)

func main() {

	fmt.Println("Selection sort: start")

	Array := []int { 4, 9, 1, 3, 1, 5, 7, 0, 11, 2, 5, 12, 10 };
	fmt.Println(Array)

	for i := 1; i < len(Array); i++ {
		j := i

		for j > 0 && Array[j-1] > Array[j] {
			Array[j-1], Array[j] = Array[j], Array[j-1]
			j--
		}
	}

	fmt.Println(Array)
	fmt.Println("Selection sort: end")
}