package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs28(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) || left.Val != right.Val {
		return false
	}
	return dfs28(left.Left, right.Right) && dfs28(left.Right, right.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs28(root.Left, root.Right)
}
