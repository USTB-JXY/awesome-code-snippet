package leetcode

// 解法一，暴力打表法
func totalNQueens(n int) int {
	res := []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352, 724}
	return res[n]
}

// 解法二，DFS 回溯法
func totalNQueens1(n int) int {
	col, dia1, dia2, row, res := make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), []int{}, 0
	putQueen52(n, 0, &col, &dia1, &dia2, &row, &res)
	return res
}

// 尝试在一个n皇后问题中, 摆放第index行的皇后位置
func putQueen52(n, index int, col, dia1, dia2 *[]bool, row *[]int, res *int) {
	if index == n {
		*res++
		return
	}
	for i := 0; i < n; i++ {
		// 尝试将第index行的皇后摆放在第i列
		if !(*col)[i] && !(*dia1)[index+i] && !(*dia2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i] = true
			(*dia1)[index+i] = true
			(*dia2)[index-i+n-1] = true
			putQueen52(n, index+1, col, dia1, dia2, row, res)
			(*col)[i] = false
			(*dia1)[index+i] = false
			(*dia2)[index-i+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
	return
}

// 解法三 二进制位操作法
// class Solution {
// 	public:
// 		int totalNQueens(int n) {
// 			int ans=0;
// 			int row=0,leftDiagonal=0,rightDiagonal=0;
// 			int bit=(1<<n)-1;//to clear high bits of the 32-bit int
// 			Queens(bit,row,leftDiagonal,rightDiagonal,ans);
// 			return ans;
// 		}
// 		void Queens(int bit,int row,int leftDiagonal,int rightDiagonal,int &ans){
// 			int cur=(~(row|leftDiagonal|rightDiagonal))&bit;//possible place for this queen
// 			if (!cur) return;//no pos for this queen
// 			while(cur){
// 				int curPos=(cur&(~cur + 1))&bit;//choose possible place in the right
// 				//update row,ld and rd
// 				row+=curPos;
// 				if (row==bit) ans++;//last row
// 				else Queens(bit,row,((leftDiagonal|curPos)<<1)&bit,((rightDiagonal|curPos)>>1)&bit,ans);
// 				cur-=curPos;//for next possible place
// 				row-=curPos;//reset row
// 			}
// 		}
// };

func mytotalNQueens(n int) int {
	visitedx, visitedxy, res, l := make([]bool, n), make([]bool, n), 0, 0

	bt(&visitedx, &visitedxy, &res, l, n)
	return res
}

func bt(vx *[]bool, vy *[]bool, res *int, l, n int) {
	if l == n {
		*res++
		return
	}
	for i := l; i < n; i++ {
		for j := 0; j < n; j++ {
			if !(*vx)[i] && !(*vy)[j] {

			}
		}
	}
}

// func solveNQueens(n int) [][]string {
// 	var res [][]string
// 	chessboard := make([][]string, n)
// 	for i := 0; i < n; i++ {
// 		chessboard[i] = make([]string, n)
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			chessboard[i][j] = "."
// 		}
// 	}
// 	var backtrack func(int)
// 	backtrack = func(row int) {
// 		if row == n {
// 			temp := make([]string, n)
// 			for i, rowStr := range chessboard {
// 				temp[i] = strings.Join(rowStr, "")
// 			}
// 			res = append(res, temp)
// 			return
// 		}
// 		for i := 0; i < n; i++ {
// 			if isValid(n, row, i, chessboard) {
// 				chessboard[row][i] = "Q"
// 				backtrack(row + 1)
// 				chessboard[row][i] = "."
// 			}
// 		}
// 	}
// 	backtrack(0)
// 	return res
// }

// func isValid(n, row, col int, chessboard [][]string) bool {
// 	for i := 0; i < row; i++ {
// 		if (*visitd)[i][col] {
// 			return false
// 		}
// 	}
// 	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
// 		if chessboard[i][j] == "Q" {
// 			return false
// 		}
// 	}
// 	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
// 		if chessboard[i][j] == "Q" {
// 			return false
// 		}
// 	}
// 	return true
// }

func isValidPos(row, col int, visitd *[][]bool) bool {
	for i := 0; i < row; i++ {
		if (*visitd)[i][col] {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*visitd)[i][j] {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len((*visitd)[0]); i, j = i-1, j+1 {
		if (*visitd)[i][j] {
			return false
		}
	}
	return true
}
func mytotalNQueens1(n int) int {
	if n == 1 {
		return 1
	}
	res := [][]string{}
	temp := []string{}
	visitd := make([][]bool, n)
	for i := 0; i < n; i++ {
		visitd[i] = make([]bool, n)
	}
	val := 0
	btnq(&res, n, 0, &temp, &visitd, &val)
	//fmt.Println(res)
	return val
}

func btnq(res *[][]string, n, index int, temp *[]string, visitd *[][]bool, val *int) {
	// fmt.Println(*temp)
	// fmt.Println("-------------")
	if len(*temp) == n {
		slice2 := make([]string, n)
		copy(slice2, *temp)
		*res = append(*res, slice2)
		*val++
		return
	}
	for i := index; i < n; i++ {
		for j := 0; j < n; j++ {
			if isValidPos(i, j, visitd) {
				*temp = append(*temp, gsting(n, j))
				(*visitd)[i][j] = true
				btnq(res, n, i+1, temp, visitd, val)
				*temp = (*temp)[:len(*temp)-1]
				(*visitd)[i][j] = false
			}
		}
	}
}
func gsting(n, index int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = '.'
	}
	b[index] = 'Q'
	return string(b)
}
