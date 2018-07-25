package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data   = []int{10, 76, 23, 35, 4, 24, 45}
	rsData = []int{4, 10, 23, 24, 35, 45, 76}
)

func TestBubbleSort(t *testing.T) {
	sortData := BubbleSort(data)
	assert.Equal(t, sortData, rsData, "they should be equal")
}
