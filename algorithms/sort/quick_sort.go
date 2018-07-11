package sort

func QuickSort(data []int) []int {
	total := len(data)
	if total <= 1 {
		return data
	}

	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < total; i++ {
		if data[i] < data[0] {
			left = append(left, data[i])
		} else {
			right = append(right, data[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)
	return arrayMerge(left, data[0], right)
}

func arrayMerge(left []int, key int, right []int) []int {
	result := make([]int, 0)
	for _, l := range left {
		result = append(result, l)
	}
	result = append(result, key)
	for _, r := range right {
		result = append(result, r)
	}
	return result
}
