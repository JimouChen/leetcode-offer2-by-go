package solution

func cuttingRope(n int) int {
	// 贪心
	if n < 4 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	return res * n
}
