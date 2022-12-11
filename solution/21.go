package solution

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		left, right := nums[i]%2, nums[j]%2
		if left == 0 && right == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else if left == 0 && right == 0 {
			j--
		} else if left == 1 && right == 0 {
			i++
			j--
		} else {
			i++
		}
	}
	return nums
}
