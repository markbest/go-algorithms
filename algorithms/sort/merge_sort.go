package sort

func MergeSort(data []int) []int {
	total := len(data)
	if total <= 1 {
		return data
	}

	left := make([]int, 0)
	right := make([]int, 0)
	mid := total / 2

	for i, x := range data {
		switch {
		case i < mid:
			left = append(left, x)
		case i >= mid:
			right = append(right, x)
		}
	}
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}
