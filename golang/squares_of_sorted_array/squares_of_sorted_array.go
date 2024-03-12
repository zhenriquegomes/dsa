package squaresofsortedarray

import "sort"

func SortedSquares(nums []int) []int {
	var squares []int
	for _, num := range nums {
		square := num * num
		squares = append(squares, square)
	}
	sort.IntSlice(squares).Sort()
	return squares
}
