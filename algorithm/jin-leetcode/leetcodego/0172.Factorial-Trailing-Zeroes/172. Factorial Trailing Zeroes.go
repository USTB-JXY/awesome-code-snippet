package leetcode

func trailingZeroes(n int) int {
	if n/5 == 0 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}
func mytrailingZeroes(n int) int {
	if n < 5 {
		return 0
	}
	return n/5 + mytrailingZeroes(n/5)
}
