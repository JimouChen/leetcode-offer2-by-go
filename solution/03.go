package solution

func findRepeatNumber(nums []int) int {
	hs := make(map[int]int, 0)
	for _, num := range nums {
		if hs[num] == 0 {
			hs[num]++
		} else {
			return num
		}
	}
	return 0
}
