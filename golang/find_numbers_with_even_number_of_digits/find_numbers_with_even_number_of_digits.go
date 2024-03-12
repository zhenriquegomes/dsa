package findnumberswithevennumberofdigits

func FindNumbers(nums []int) int {
	var even_number_of_digits int
	for _, num := range nums {
		var (
			digits int
		)
		for {
			digits++
			if (num / 10) == 0 {
				break
			}
			num /= 10
		}
		if (digits % 2) == 0 {
			even_number_of_digits++
		}
	}
	return even_number_of_digits
}
