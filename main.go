package main

import (
	"fmt"
	"package/chapter8"
	"package/util"
)

func main() {
	util.PrintUtil("hola")
	x := []int{1, 10, 3, 4, 5, 6, 7, 8}
	chapter8.SortedFromSmallesToLargestOptimizated(x)
	fmt.Println("...")
	chapter8.SortedFromSmallesToLargest(x)
	//capitulo8.SortedFromLargestToSmalles(x)

}
