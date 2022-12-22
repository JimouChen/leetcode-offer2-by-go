package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	res  [][]int
	path []int
)

func dfs34(root *TreeNode, target int) {
	if root != nil {
		path = append(path, root.Val)
		target -= root.Val
		//defer func() {path = path[:len(path)-1]}()
		if target == 0 && root.Left == nil && root.Right == nil {
			res = append(res, append([]int(nil), path...))
		}
		dfs34(root.Left, target)
		dfs34(root.Right, target)
		path = path[:len(path)-1]
	}
}

func pathSum(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	path = make([]int, 0)
	dfs34(root, target)
	return res
}
