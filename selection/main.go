package main

import (
	"fmt"
)

func main() {

	fmt.Println("Selection sort: start")

	Array := []int { 4, 9, 3, 1, 7, 0, 11, 2, 5, 12, 10 };
	fmt.Println(Array)

	for j := 0; j < len(Array) - 1; j++ {

		min := j

		for i := j + 1; i < len(Array); i++ {
			if Array[i] < Array[min] {
				min = i
			}
		}

		if min != j {
			Array[j], Array[min] = Array[min], Array[j]
		}

	}

	fmt.Println(Array)
	fmt.Println("Selection sort: end")
}