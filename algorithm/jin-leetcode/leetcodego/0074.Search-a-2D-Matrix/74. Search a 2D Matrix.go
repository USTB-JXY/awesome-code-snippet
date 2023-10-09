package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	for low <= high {
		mid := low + (high-low)>>1
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
func mysearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid/n][mid%n] == target {
			return true
		}
		if matrix[mid/n][mid%n] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
