package findmaxconsecutiveones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Arr                  []int
	ExpectedNumberOfOnes int
}

func TestGivenAnArray_WhenFindingMaxConsecutiveOnes_ShouldReturnTheExpectedResult(t *testing.T) {
	testCases := []TestCase{
		{Arr: []int{1, 1, 0, 1, 1, 1}, ExpectedNumberOfOnes: 3},
		{Arr: []int{1, 1, 1, 1}, ExpectedNumberOfOnes: 4},
		{Arr: []int{1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1}, ExpectedNumberOfOnes: 6},
		{Arr: []int{0, 0, 0}, ExpectedNumberOfOnes: 0},
		{Arr: []int{}, ExpectedNumberOfOnes: 0},
	}
	for _, testCase := range testCases {
		numberOfOnes := findMaxConsecutiveOnes(testCase.Arr)
		assert.Equal(t, testCase.ExpectedNumberOfOnes, numberOfOnes)
	}
}
