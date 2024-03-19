package findmaxconsecutiveones

func findMaxConsecutiveOnes(nums []int) int {
	var (
		maxConsecutiveOnes int
		onesCount          int
	)
	for _, num := range nums {
		if num == 1 {
			onesCount++
			if onesCount > maxConsecutiveOnes {
				maxConsecutiveOnes = onesCount
			}
		} else {
			onesCount = 0
		}
	}
	return maxConsecutiveOnes
}
