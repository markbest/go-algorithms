package sort

func SelectSort(data []int) []int {
	total := len(data)
	if total == 0 {
		return nil
	}

	var min, minKey int
	for i := 0; i < total-1; i++ {
		min = data[i]
		minKey = i
		for j := i + 1; j < total; j++ {
			if min > data[j] {
				min = data[j]
				minKey = j
			}
		}
		data[i], data[minKey] = data[minKey], data[i]
	}
	return data
}
