package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	HeapSort(data)
	assert.Equal(t, data, rsData, "they should be equal")
}
