package solution

import (
	"sort"
)

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curRight := res[len(res)-1][1]
		if intervals[i][0] <= curRight {
			res[len(res)-1][1] = max(intervals[i][1], curRight)
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
