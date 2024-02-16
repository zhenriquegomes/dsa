package binarysearch

func BinarySearch(nums []int, target int) int {
	begin := 0
	end := len(nums) - 1
	for begin <= end {
		mid := (begin+end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			begin = mid + 1
		} else {
			return mid
		}
	}
	return -1
}