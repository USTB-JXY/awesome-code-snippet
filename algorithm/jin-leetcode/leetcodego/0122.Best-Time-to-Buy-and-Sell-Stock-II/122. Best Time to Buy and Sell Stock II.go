package leetcode

func maxProfit122(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

func sel2(nums []int) int {
	if len(nums) < 0 {
		return 0
	}
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			max += nums[i+1] - nums[i]
		}
	}
	return max
}
