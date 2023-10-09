package leetcode

var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	res, visited := 0, make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				searchIslands(grid, &visited, i, j)
				res++
			}
		}
	}
	return res
}

func searchIslands(grid [][]byte, visited *[][]bool, x, y int) {
	(*visited)[x][y] = true
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		if isInBoard(grid, nx, ny) && !(*visited)[nx][ny] && grid[nx][ny] == '1' {
			searchIslands(grid, visited, nx, ny)
		}
	}
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func mynumIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	res, visit := 0, make([][]bool, m)
	//now visit is [[] [] [] []]
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	//now visit is [[false false false false false] [false false false false false] ......]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visit[i][j] {
				dfs(i, j, grid, &visit)
				res++
			}
		}
	}
	return res
}
func dfs(i, j int, grid [][]byte, visit *[][]bool) {
	(*visit)[i][j] = true
	for k := 0; k < 4; k++ {
		x, y := i+dir[k][0], j+dir[k][1]
		if x > -1 && y > -1 && y < len(grid[0]) && x < len(grid) {
			if grid[x][y] == '1' && !((*visit)[x][y]) {
				dfs(x, y, grid, visit)
			}
		}
	}
}
