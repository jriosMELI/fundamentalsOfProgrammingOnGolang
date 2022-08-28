package chapter8

import "fmt"

func SortedFromSmallesToLargestBubbleSort(arr []int) {
	for i := 0; i < (len(arr) - 1); i++ {
		for j := 0; j < (len(arr) - 1); j++ {
			if arr[j] > arr[j+1] {
				aux := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = aux
			}
		}
		fmt.Println(arr)
	}
}

func SortedFromLargestToSmallesBubbleSort(arr []int) {
	for i := 0; i < (len(arr) - 1); i++ {
		for j := 0; j < (len(arr) - 1); j++ {
			if arr[j] < arr[j+1] {
				aux := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = aux
			}
		}
	}
	fmt.Println(arr)
}

func SortedFromSmallesToLargestOptimizatedBubbleSort(arr []int) {
	sorted := false
	i := 0

	for !sorted && i <= len(arr)-1 {
		sorted = true
		for j := 0; j < (len(arr) - 1); j++ {
			if arr[j] > arr[j+1] {
				aux := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = aux
				sorted = false
			}
		}
		fmt.Println(arr)
	}
}
