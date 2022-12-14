package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	if root != nil {
		left := mirrorTree(root.Left)
		right := mirrorTree(root.Right)
		root.Left = right
		root.Right = left
		return root
	}
	return nil
}
