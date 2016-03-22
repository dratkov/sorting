package main

import (
	"fmt"
)

func main() {

	fmt.Println("Merge sort: start")

	Array := []int { 4, 9, 1, 3, 1, 5, 14, 7, 0, 11, 2, 5, 12, 10 };
	fmt.Println(Array)

	Array2 := merge_sort(Array)

	fmt.Println(Array2)
	fmt.Println("Merge sort: end")

}

func merge_sort( a []int ) []int {
	if len(a) == 1 {
		return a
	}

	var left, right []int
	for i := 0; i < len(a); i++ {
		if i % 2 == 0 {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}

	left = merge_sort(left)
	right = merge_sort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	l, r := 0, 0

	for len(left) + len(right) != len(result) {
		switch {
			case l == len(left):
				result = append(result, right[r:]...)
			case r == len(right):
				result = append(result, left[l:]...)
			case left[l] > right[r]:
				result = append(result, right[r])
				r++
			default:
				result = append(result, left[l])
				l++
		}
	}

	return result
}
