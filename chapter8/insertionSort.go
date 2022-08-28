package chapter8

import "fmt"

func SortedFromSmallesToLargestInsertionSort(arr []int) {
	for j := 1; j < (len(arr)); j++ {
		position := j
		auxiliar := arr[j]
		for position > 0 && arr[position-1] > auxiliar {
			arr[position] = arr[position-1]
			position = position - 1
		}
		arr[position] = auxiliar
	}
	fmt.Println(arr)
}
