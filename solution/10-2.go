package solution

func numWays(n int) int {
	// 1 1 2 3
	if n <= 1 {
		return 1
	}
	cur := 1
	next := 1
	for i := 2; i <= n; i++ {
		temp := cur
		cur = next
		next = (next + temp) % (1e9 + 7)
	}
	return next
}
