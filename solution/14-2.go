package solution

func cuttingRope2(n int) int {
	if n < 4 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res = res * 3 % (1e9 + 7)
		n -= 3
	}
	return res * n % (1e9 + 7)
}
