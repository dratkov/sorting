package main

import (
	"fmt"
)

func quick(data []int) {
	if len(data) < 2 {
		return
	}

	pivot := data[0]
	l, r := 1, len(data) - 1
	for l <= r {
		for l <= r && data[l] <= pivot {
			l++
		}

		for r >= l && data[r] >= pivot {
			r--
		}

		if l < r {
			data[r], data[l] = data[l], data[r]
		}
	}
	if r > 0 {
		data[0], data[r] = data[r], data[0]
		quick(data[0:r])
	}
	quick(data[l:])
}

func main() {

	fmt.Println("Quick sort: start")

	Array := []int { 7, 1, 2, 300, 10, 8, 9 };
	fmt.Println(Array)

	quick(Array)

	fmt.Println(Array)
	fmt.Println("Quick sort: end")
}