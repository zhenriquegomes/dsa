package squaresofsortedarray

import "sort"

func sortedSquares(nums []int) []int {
	squares := make([]int, len(nums))
	for i, num := range nums {
		square := num * num
		squares[i] = square
	}
	sort.IntSlice(squares).Sort()
	return squares
}
