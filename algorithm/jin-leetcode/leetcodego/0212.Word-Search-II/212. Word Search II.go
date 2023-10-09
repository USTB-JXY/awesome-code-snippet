package leetcode

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	for _, v := range words {
		if exist(board, v) {
			res = append(res, v)
		}
	}
	return res
}

// these is 79 solution
var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false
}
func myfindWords(board [][]byte, words []string) []string {
	res := []string{}
	for _, v := range words {
		if myexist(board, v) {
			res = append(res, v)
		}
	}
	return res
}
func myexist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited[i][j] = true
				if bt(i, j, 0, word, &visited, board) {
					return true
				} else {
					visited[i][j] = false
				}
			}
		}
	}
	return false
}

func bt(i, j, l int, word string, visited *[][]bool, board [][]byte) bool {
	// fmt.Println(word[:l+1])
	// fmt.Println(*visited)
	if l == len(word)-1 {
		return true
	}
	for k := 0; k < 4; k++ {
		x, y := i+dir[k][0], j+dir[k][1]
		if x >= 0 && y >= 0 && x < len(board) && y < len(board[0]) {
			//fmt.Println(x, y)
			if !(*visited)[x][y] && board[x][y] == word[l+1] {
				(*visited)[x][y] = true
				if bt(x, y, l+1, word, visited, board) {
					return true
				} else {
					(*visited)[x][y] = false
				}

			}
		}
	}
	return false
}
