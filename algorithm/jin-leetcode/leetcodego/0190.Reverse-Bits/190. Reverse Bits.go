package leetcode

import "fmt"

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32 && num > 0; i++ {
		res = res<<1 | num&1
		num >>= 1
		fmt.Printf("【res】:%b  【num】:%b\n", res, num)
	}
	return res
}

func myreverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | num&1
		num = num >> 1
	}
	return res
}
