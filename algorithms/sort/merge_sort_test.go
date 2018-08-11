package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	sortData := MergeSort(data)
	assert.Equal(t, sortData, rsData, "they should be equal")
}
