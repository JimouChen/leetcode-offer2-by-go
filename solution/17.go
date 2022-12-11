package solution

import "math"

func printNumbers(n int) []int {
	length := int(math.Pow(10, float64(n))) - 1
	res := make([]int, length)
	for i := 1; i <= length; i++ {
		res[i-1] = i
	}
	return res
}
