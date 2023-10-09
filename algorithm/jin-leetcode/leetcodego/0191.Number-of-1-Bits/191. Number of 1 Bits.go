package leetcode

import "math/bits"

// 解法一
func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

// 解法二
func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
func myhammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

//1,000 001111101000
//999   001111100111
// &    001111100000

// 1010100000100001
// 1010100000100000
func myhammingWeight2(num uint32) int {
	res := 0
	for num != 0 {
		num &= num - 1
		res++
	}
	return res
}
