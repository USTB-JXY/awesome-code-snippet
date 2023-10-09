package leetcode

func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}

// Source : https://oj.leetcode.com/problems/trapping-rain-water/
// Author : Hao Chen
// Date   : 2014-07-02

/**********************************************************************************
*
* Given n non-negative integers representing an elevation map where the width of each bar is 1,
* compute how much water it is able to trap after raining.
*
* For example,
* Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
*
*     ^
*     |
*   3 |                       +--+
*     |                       |  |
*   2 |          +--+xxxxxxxxx|  +--+xx+--+
*     |          |  |xxxxxxxxx|  |  |xx|  |
*   1 |   +--+xxx|  +--+xxx+--+  |  +--+  +--+
*     |   |  |xxx|  |  |xxx|  |  |  |  |  |  |
*   0 +---+--+---+--+--+---+--+--+--+--+--+--+----->
*       0  1   0  2  1   0  1  3  2  1  2  1
*
* The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case,
* 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
*
**********************************************************************************/
func mytrap(height []int) int {
	if len(height) < 1 {
		return 0
	}
	res, left, right := 0, 0, len(height)-1

	maxleft, maxright := 0, 0
	for left <= right {
		if height[left] < height[right] {
			if height[left] > maxleft {
				maxleft = height[left]
			} else {
				res += maxleft - height[left]
			}
			left++
		} else {
			if height[right] > maxright {
				maxright = height[right]
			} else {
				res += maxright - height[right]
			}
			right--
		}
	}
	return res
}
