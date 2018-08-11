package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSort(t *testing.T) {
	sortData := ShellSort(data)
	assert.Equal(t, sortData, rsData, "they should be equle")
}
