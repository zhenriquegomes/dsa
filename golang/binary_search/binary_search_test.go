package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Arr                 []int
	Target              int
	ExpectedTargetIndex int
}

func TestGivenAnOrderedArray_WhenSearchingAnTarget_ThenShouldReturnTargetIndex(t *testing.T) {
	arr := []int{-1, 0, 4, 7, 8, 15}
	testCases := []TestCase{
		{Arr: arr, Target: 4, ExpectedTargetIndex: 2},
		{Arr: arr, Target: -1, ExpectedTargetIndex: 0},
		{Arr: arr, Target: 10, ExpectedTargetIndex: -1},
	}
	for _, testCase := range testCases {
		target_index := binarySearch(testCase.Arr, testCase.Target)
		assert.Equal(t, testCase.ExpectedTargetIndex, target_index)
	}
}
