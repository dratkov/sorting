package main

import (
	"fmt"
)

func sort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		heapify(a[0:i+1])
		a[0], a[i] = a[i], a[0]
	}
}

func heapify(a []int) {
	length := len(a)
	for i := 0; i < length; i++ {
		heapify_left(a, i)
		heapify_right(a, i)
	}
}

func heapify_left(a []int, parent int) {
	child := parent * 2 + 1
	for parent >= 0 && child < len(a) && a[child] > a[parent] {
		a[parent], a[child] = a[child], a[parent]
		heapify_left(a, ( parent - 1 ) / 2 )
	}
}

func heapify_right(a []int, parent int) {
	child := parent * 2 + 2
	for parent >= 0 && child < len(a) && a[child] > a[parent] {
		a[parent], a[child] = a[child], a[parent]
		heapify_left(a, ( parent - 2 ) / 2 )
	}
}

func main() {

	fmt.Println("Heap sort: start")

	Array := []int { 4, 80, 9, 1, 3, 22, 1, 5, 0, 14, 1, 7, 11, 0, 2, 5, 12, 10 };
	fmt.Println(Array)

	sort(Array)

	fmt.Println(Array)
	fmt.Println("Heap sort: end")
}