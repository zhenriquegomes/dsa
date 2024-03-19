package duplicatezeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Arr         []int
	ExpectedArr []int
}

func TestGivenAnArray_WhenDuplicatingZeros_ShouldShiftTheElementsOnTheRight(t *testing.T) {
	testCases := []TestCase{
		{Arr: []int{1, 0, 2, 3, 0, 4, 5, 0}, ExpectedArr: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{Arr: []int{1, 2, 3, 4, 5}, ExpectedArr: []int{1, 2, 3, 4, 5}},
		{Arr: []int{3, 0, 4, 5, 0}, ExpectedArr: []int{3, 0, 0, 4, 5}},
	}
	for _, testCase := range testCases {
		duplicateZeros(testCase.Arr)
		assert.Equal(t, testCase.Arr, testCase.ExpectedArr)
	}

}
