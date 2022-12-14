package solution

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	j := 0
	for _, num := range pushed {
		stack = append(stack, num)
		for len(stack) > 0 && (popped[j] == stack[len(stack)-1]) {
			stack = append(stack[:len(stack)-1])
			j++
		}
	}
	return len(stack) == 0
}
