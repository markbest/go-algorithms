package search

import "math"

func BinarySearch(data []int, v int) int {
	left, right, mid := 1, len(data), 0
	for {
		mid = int(math.Floor(float64((left + right) / 2)))
		if data[mid] > v {
			right = mid - 1
		} else if data[mid] < v {
			left = mid + 1
		} else {
			break
		}

		if left > right {
			mid = -1
			break
		}
	}
	return mid
}
