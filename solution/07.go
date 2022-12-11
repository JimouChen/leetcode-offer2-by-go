package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findIndex(a []int, target int) int {
	for idx, v := range a {
		if v == target {
			return idx
		}
	}
	return -1
}

func dfs(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	preRootVal := preorder[0]
	inRootIdx := findIndex(inorder, preRootVal)
	root := &TreeNode{Val: preRootVal}
	// inRootIdx也是表示左子树的个数
	root.Left = dfs(preorder[1:1+inRootIdx], inorder[:inRootIdx])
	root.Right = dfs(preorder[inRootIdx+1:], inorder[inRootIdx+1:])
	return root
}

// BuildTree 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
func buildTree(preorder []int, inorder []int) *TreeNode {
	index := make(map[int]int)
	for idx, v := range inorder {
		index[v] = idx
	}
	return dfs(preorder, inorder)
}
