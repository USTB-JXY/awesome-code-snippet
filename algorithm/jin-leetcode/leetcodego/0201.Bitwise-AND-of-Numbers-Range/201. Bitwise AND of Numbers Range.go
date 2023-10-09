package leetcode

import "fmt"

// 解法一
func rangeBitwiseAnd1(m int, n int) int {
	if m == 0 {
		return 0
	}
	moved := 0
	for m != n {
		m >>= 1
		n >>= 1
		moved++
	}
	return m << uint32(moved)
}

// 解法二 Brian Kernighan's algorithm
func rangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= (n - 1) // 清除最低位的 1
	}
	return n
}

// 11010
// 11011
// 11100
// 11101
// 11110
func myrangeBitwiseAnd(m int, n int) int {
	for n > m {
		n &= n - 1
		fmt.Printf("【res】:%b  【num】:%b\n", m, n)
	}
	return n
}
