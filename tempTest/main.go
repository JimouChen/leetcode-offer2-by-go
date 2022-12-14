package main

/*

1，主导完成1130特性详细设计及转测，带领专项小组成员保障版本高质量发布，满足商用要求：
2，pmd优化专项之瑶光智能调度联合开发，完成阶段一目标联调验证，上线目标达成。
3，完成ebackup在堆叠+政务云三网隔离场景上线场景分析验证给你一个整数数组?nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。

?

示例 1：

输入：nums = [1,2,1,3,2,5]
输出：[3,5]
解释：[5, 3] 也是有效的答案。
示例 2：

输入：nums = [-1,0]
输出：[-1,0]
示例 3：

输入：nums = [0,1]
输出：[1,0]
?

提示：

2 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
除两个只出现一次的整数外，nums 中的其他数字都出现两次

*/
import "fmt"

func singleNumber(nums []int) []int {
	hs := make(map[int]int)
	res := make([]int, 0)
	for _, num := range nums {
		hs[num]++
	}
	for k, v := range hs {
		if v == 1 {
			res = append(res, k)
		}
		if len(res) == 2 {
			return res
		}
	}

	return res
}
func main() {
	// to test the solution function
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{1, 0}))
	fmt.Println(singleNumber([]int{-1, 0}))
}
