package sort

func SelectSort(data []int) []int {
	total := len(data)
	if total == 0 {
		return nil
	}

	for i := 0; i < total-1; i++ {
		min := i
		for j := i + 1; j < total; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		if min != i {
			data[i], data[min] = data[min], data[i]
		}
	}
	return data
}
