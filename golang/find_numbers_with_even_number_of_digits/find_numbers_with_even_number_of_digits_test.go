package findnumberswithevennumberofdigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Arr           []int
	ExpectedCount int
}

func TestGivenAnArray_WhenLookingForNumbersWithEvenNumberOfDigits_ShouldReturnTheCountOfNumbersWithEvenNumberOfDigits(t *testing.T) {
	testCases := []TestCase{
		{Arr: []int{12, 345, 2, 6, 7896}, ExpectedCount: 2},
		{Arr: []int{17, 23, 35, 6, 9}, ExpectedCount: 3},
		{Arr: []int{1, 5, 256, 600, 78960}, ExpectedCount: 0},
	}
	for _, testCase := range testCases {
		count := findNumbers(testCase.Arr)
		assert.Equal(t, testCase.ExpectedCount, count)
	}
}
