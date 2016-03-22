package main

import (
	"fmt"
)

func main() {

	fmt.Println("Shell sort: start")

	Array := []int { 4, 9, 1, 3, 1, 5, 7, 0, 11, 2, 5, 12, 10 };
	fmt.Println(Array)

	for gap := len(Array) / 2; gap > 0; gap /= 2 {
		for p := gap; p < len(Array); p++ {
			for j := p; j >= gap && Array[j - gap] > Array[j]; j -=gap {
				Array[j - gap], Array[j] = Array[j], Array[j - gap]
			}
		}
	}

	fmt.Println(Array)
	fmt.Println("Shell sort: end")
}