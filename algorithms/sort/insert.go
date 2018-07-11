package sort

func InsertSort(data []int) []int {
	var key, pos int
	total := len(data)
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
