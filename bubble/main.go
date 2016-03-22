package main

import (
	"fmt"
)

func main() {

	fmt.Println("Buble sort: start")

	Array := []int { 4, 9, 3, 1, 7, 0, 2, 5, 12, 10 };
	fmt.Println(Array)
	for {
		is_was_swap := false
		for index := 0; index < len(Array) - 2; index++ {
			if Array[index] > Array[index+1] {
				Array[index], Array[index+1] = Array[index+1], Array[index]
				is_was_swap = true
			}
		}
		if is_was_swap == false {
			break
		}
	}

	fmt.Println(Array)
	fmt.Println("Bubble sort: end")

}