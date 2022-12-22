package solution

func dfs33(i, j int, postorder []int) bool {
	if i >= j {
		return true
	}
	Idx := i
	for postorder[Idx] < postorder[j] {
		Idx++
	}
	newRightIdx := Idx
	// 此时左子树为 [i: newRightIdx - 1]
	// 判断剩下的区间是否满足都大于根节点
	for postorder[Idx] > postorder[j] {
		Idx++
	}
	return Idx == j && dfs33(i, newRightIdx-1, postorder) && dfs33(newRightIdx, j-1, postorder)
}
func verifyPostorder(postorder []int) bool {
	return dfs33(0, len(postorder)-1, postorder)
}
