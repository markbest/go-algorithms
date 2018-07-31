package sort

func BubbleSort(data []int) []int {
	total := len(data)
	if total == 0 {
		return nil
	}

	for i := 0; i < total; i++ {
		for j := 0; j < total-i-1; j++ {
			if data[j] > data[j+1] {
				data[j+1], data[j] = data[j], data[j+1]
			}
		}
	}
	return data
}
