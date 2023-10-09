package leetcode

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] += candies[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	total := 0
	for _, candy := range candies {
		total += candy + 1
	}
	return total
}

func mycandy(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	val, n := 0, len(nums)
	candies := make([]int, n)
	candies[0] = 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] {
			candies[i+1] = candies[i] + 1
		} else {
			candies[i+1] = 1
		}
	}
	val += candies[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			candies[i] = candies[i+1] + 1
		}
		val += candies[i]
	}
	return val
}

// i< n-2
//i+1< n-1
