package leetcode

import (
	"fmt"
	"sort"
)

// 解法一 O(n^2) DP
func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1 10 20 15
// 解法二 O(n log n) DP
func lengthOfLIS1(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
		fmt.Println(dp)
	}
	return len(dp)
}

func mylengthOfLIS1(nums []int) int {
	dp := []int{}
	for _, v := range nums {
		index := sort.SearchInts(dp, v)
		if index == len(dp) {
			dp = append(dp, v)
		} else {
			dp[index] = v
		}
	}
	return len(dp)
}
