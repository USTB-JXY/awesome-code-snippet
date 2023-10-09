package leetcode

func minSubArrayLen(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for right, v := range nums {
		sum += v
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minSubArrayLen1(target int, nums []int) int {
	i := 0
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
func myinSubArrayLen(target int, nums []int) int {
	i, l, sum, result := 0, len(nums), 0, len(nums)+1
	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			sublength := j - i + 1
			if sublength < result {
				result = sublength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}
