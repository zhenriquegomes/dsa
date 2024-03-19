package squaresofsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Arr         []int
	ExpectedArr []int
}

func TestGivenAnArray_WhenSquaringItsNumbers_ShouldReturnAnSortedArray(t *testing.T) {
	testCases := []TestCase{
		{Arr: []int{-4, -1, 0, 3, 10}, ExpectedArr: []int{0, 1, 9, 16, 100}},
		{Arr: []int{-7, -3, 2, 3, 11}, ExpectedArr: []int{4, 9, 9, 49, 121}},
		{Arr: []int{}, ExpectedArr: []int{}},
	}
	for _, testCase := range testCases {
		squaredArray := sortedSquares(testCase.Arr)
		assert.Equal(t, testCase.ExpectedArr, squaredArray)
	}
}
