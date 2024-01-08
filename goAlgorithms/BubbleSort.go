package main

import (
	"fmt"
)

type Sorter interface {
	Sort([]int)
}

type BubbleSort struct{}

func (bs BubbleSort) Sort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	unsortedArray := []int{64, 34, 25, 12, 22, 11, 90}

	sorter := BubbleSort{}

	fmt.Println("Unsorted array:", unsortedArray)
	sorter.Sort(unsortedArray)
	fmt.Println("Sorted array:", unsortedArray)
}
