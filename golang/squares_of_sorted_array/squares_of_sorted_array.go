package squaresofsortedarray

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	squares := make([]int, 0, len(nums))
	l, r := 0, len(nums)-1
	for l <= r {
		if math.Abs(float64(nums[l])) > math.Abs(float64(nums[r])) {
			squares = append(squares, nums[l]*nums[l])
			l++
		} else {
			squares = append(squares, nums[r]*nums[r])
			r--
		}
	}
	reverseSlice(squares)
	return squares
}

func reverseSlice(nums []int) {
	fmt.Println(nums)
	i, j := 0, len(nums)-1
	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	fmt.Println(nums)
}
