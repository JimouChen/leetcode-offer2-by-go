package solution

func fib(n int) int {
	if n <= 1 {
		return n
	}
	cur := 0
	next := 1
	for i := 2; i <= n; i++ {
		temp := cur
		cur = next
		next = (temp + next) % (1e9 + 7)
	}
	return next
}
