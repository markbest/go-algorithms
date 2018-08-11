package search

func BinarySearch(data []int, v int) int {
	left, right, mid := 0, len(data)-1, 0
	for {
		mid = left + (right-left)/2
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
