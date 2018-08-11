package sort

func HeapSort(values []int) {
	buildHeap(values)
	for i := len(values); i > 1; i-- {
		values[0], values[i-1] = values[i-1], values[0]
		adjustHeap(values[:i-1], 0)
	}
}

// 初始构造堆
func buildHeap(values []int) {
	for i := len(values); i >= 0; i-- {
		adjustHeap(values, i)
	}
}

// 调整堆
func adjustHeap(values []int, pos int) {
	node := pos
	length := len(values)
	for node < length {
		var child = 0
		if 2*node+2 < length {
			if values[2*node+1] > values[2*node+2] {
				child = 2*node + 1
			} else {
				child = 2*node + 2
			}
		} else if 2*node+1 < length {
			child = 2*node + 1
		}

		if child > 0 && values[child] > values[node] {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}
	}
}
