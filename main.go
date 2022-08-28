package main

import (
	"package/chapter8"
	"package/util"
)

func main() {
	util.PrintUtil("hola")
	x := []int{14, 10, 3, 4, 5, 6, 7, 1}
	/*
		// Example Bubble Sort
			chapter8.SortedFromSmallesToLargestOptimizatedBubbleSort(x)
			chapter8.SortedFromSmallesToLargestBubbleSort(x)
			chapter8.SortedFromLargestToSmallesBubbleSort(x)
	*/
	// Example Insertion Sort
	chapter8.SortedFromSmallesToLargestInsertionSort(x)
}
