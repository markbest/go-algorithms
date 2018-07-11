package main

import "fmt"

func BubbleSort(data []int) []int {
	total := len(data)
	for i := 0; i < total; i++ {
		for j := total - 1; j > i; j-- {
			if data[j] < data[j-1] {
				tmp := data[j]
				data[j] = data[j-1]
				data[j-1] = tmp
			}
		}
	}
	return data
}

func BubbleSort1(data []int) []int {
	total := len(data)
	for i := 0; i < total-1; i++ {
		for j := 0; j < total-1-i; j++ {
			if data[j] > data[j+1] {
				tmp := data[j+1]
				data[j+1] = data[j]
				data[j] = tmp
			}
		}
	}
	return data
}

func main() {
	data := []int{10, 76, 23, 35, 4, 24, 45}
	sortData := BubbleSort(data)
	fmt.Println(sortData)
	sortData = BubbleSort1(data)
	fmt.Println(sortData)
}
