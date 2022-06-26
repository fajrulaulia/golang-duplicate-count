package golangdulicatecount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDuplicateCheckerValid(t *testing.T) {
	assert.Equal(t, 2, DuplicateCount([]int{1, 2, 3, 3, 4, 5, 6, 2, 9}), "asessrtion 1")
	assert.Equal(t, 2, DuplicateCount([]int{5, 2, 3, 4, 5, 2, 5, 7, 5}), "asessrtion 2")
	assert.Equal(t, 1, DuplicateCount([]int{1, 3, 4, 3, 5, 6, 7, 8, 9, 10, 11}), "asessrtion 3")
	assert.Equal(t, 6, DuplicateCount([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 8, 6, 4, 1, 2}), "asessrtion 4")
}

func TestCountDuplicateCheckerInValid(t *testing.T) {
	assert.NotEqual(t, 3, DuplicateCount([]int{1, 2, 3, 3, 4, 5, 6, 2, 9}), "asessrtion 1")
	assert.NotEqual(t, 4, DuplicateCount([]int{5, 2, 3, 4, 5, 2, 5, 7, 5}), "asessrtion 2")
	assert.NotEqual(t, 3, DuplicateCount([]int{1, 3, 4, 3, 5, 6, 7, 8, 9, 10, 11}), "asessrtion 3")
	assert.NotEqual(t, 8, DuplicateCount([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 8, 6, 4, 1, 2}), "asessrtion 4")
}
