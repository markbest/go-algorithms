package sort

func ShellSort(data []int) []int {
	total := len(data)
	if total == 0 {
		return nil
	}

	gap := 2
	step := total / gap
	for step >= 1 {
		for i := step; i < len(data); i++ {
			for j := i; j >= step && data[j] < data[j-step]; j -= step {
				data[j], data[j-step] = data[j-step], data[j]
			}
		}
		step = step / gap
	}
	return data
}
