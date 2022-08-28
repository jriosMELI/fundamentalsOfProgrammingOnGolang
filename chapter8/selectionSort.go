package chapter8

import "fmt"

func SortedFromSmallesToLargestSelectionSort(arr []int) {
	for i := 0; i < (len(arr) - 1); i++ {
		minimun := i
		for j := i + 1; j < (len(arr)); j++ {
			if arr[j] < arr[minimun] {
				minimun = j
			}
		}
		auxiliar := arr[i]
		arr[i] = arr[minimun]
		arr[minimun] = auxiliar
	}
	fmt.Println(arr)
}
