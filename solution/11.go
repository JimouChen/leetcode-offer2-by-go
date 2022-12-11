package solution

import "sort"

func minArray(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}
