package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	sortData := QuickSort(data)
	assert.Equal(t, sortData, rsData, "they should be equal")
}
