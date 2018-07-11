package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	sortData := InsertSort(data)
	assert.Equal(t, sortData, rsData, "they should be equal")
}
