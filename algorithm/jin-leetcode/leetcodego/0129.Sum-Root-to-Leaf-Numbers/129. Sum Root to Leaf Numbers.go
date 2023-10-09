package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	res := 0
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += sum
		return
	}
	dfs(root.Left, sum, res)
	dfs(root.Right, sum, res)
}

// Input: [1,2,3]
// 1
// / \
// 2   3
// Output: 25
// Explanation:
// The root-to-leaf path 1->2 represents the number 12.
// The root-to-leaf path 1->3 represents the number 13.
// Therefore, sum = 12 + 13 = 25.

// **Example 2:**

// Input: [4,9,0,5,1]
// 4
// / \
// 9   0
// / \
// 5   1
// Output: 1026
// Explanation:
// The root-to-leaf path 4->9->5 represents the number 495.
// The root-to-leaf path 4->9->1 represents the number 491.
// The root-to-leaf path 4->0 represents the number 40.
// Therefore, sum = 495 + 491 + 40 = 1026.

func mysumNumbers(root *TreeNode) int {
	return mydfs(root, 0) / 2
}
func mydfs(root *TreeNode, res int) int {
	if root == nil {
		return res
	}
	return mydfs(root.Left, res*10+root.Val) + mydfs(root.Right, res*10+root.Val)
}
