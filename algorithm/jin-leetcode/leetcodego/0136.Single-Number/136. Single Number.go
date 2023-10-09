package leetcode

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func mysingleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
