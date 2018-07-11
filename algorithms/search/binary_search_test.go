package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	data   = []int{1, 2, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	rsData = 6
)

func TestBinarySearch(t *testing.T) {
	rs := BinarySearch(data, 30)
	assert.Equal(t, rs, rsData, "they should be equal")
}