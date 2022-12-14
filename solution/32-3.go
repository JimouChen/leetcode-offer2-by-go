package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	idx := 0
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		temp := make([]int, 0)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = append(queue[1:])
			temp = append(temp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		if idx%2 == 1 {
			temp = reverseSlice(temp)
		}
		res = append(res, temp)
		idx++
	}

	return res
}

func reverseSlice(a []int) []int {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
	return a
}
