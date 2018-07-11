package sort

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
