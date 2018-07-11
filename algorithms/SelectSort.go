package main

import "fmt"

func SelectSort(array []int) []int {
	var tmp, min, minKey int
	total := len(array)
	for i := 0; i < total-1; i++ {
		min = array[i]
		minKey = i
		for j := i + 1; j < total; j++ {
			if min > array[j] {
				min = array[j]
				minKey = j
			}
		}

		tmp = array[i]
		array[i] = array[minKey]
		array[minKey] = tmp
	}
	return array
}

func main() {
	data := []int{10, 76, 23, 35, 4, 24, 45}
	sortData := SelectSort(data)
	fmt.Println(sortData)
}
