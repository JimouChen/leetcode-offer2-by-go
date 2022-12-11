package solution

func dfs12(i, j, k int, board [][]byte, word string) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}

	temp := board[i][j]
	board[i][j] = '.'
	res := dfs12(i+1, j, k+1, board, word) ||
		dfs12(i, j+1, k+1, board, word) ||
		dfs12(i-1, j, k+1, board, word) ||
		dfs12(i, j-1, k+1, board, word)
	board[i][j] = temp
	return res

}
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs12(i, j, 0, board, word) {
				return true
			}
		}
	}
	return false
}
