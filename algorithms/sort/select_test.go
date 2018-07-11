package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {
	sortData := SelectSort(data)
	assert.Equal(t, sortData, rsData, "they should be equal")
}
