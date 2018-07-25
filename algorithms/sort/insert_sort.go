package sort

func InsertSort(data []int) []int {
	total := len(data)
	if total == 0 {
		return nil
	}

	var key, pos int
	for i := 0; i < total; i++ {
		key = data[i]
		pos = i

		for pos > 0 && data[pos-1] > key {
			data[pos] = data[pos-1]
			pos = pos - 1
		}
		data[pos] = key
	}
	return data
}
