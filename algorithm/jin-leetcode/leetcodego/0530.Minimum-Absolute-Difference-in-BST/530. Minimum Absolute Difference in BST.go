package leetcode

import (
	"math"

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

func getMinimumDifference(root *TreeNode) int {
	res, nodes := math.MaxInt16, -1
	dfsBST(root, &res, &nodes)
	return res
}

func dfsBST(root *TreeNode, res, pre *int) {
	if root == nil {
		return
	}
	dfsBST(root.Left, res, pre)
	if *pre != -1 {
		*res = min(*res, abs(root.Val-*pre))
	}
	*pre = root.Val
	dfsBST(root.Right, res, pre)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func mygetMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt16, -1
	dfs(root, &res, &pre)
	return res
}
func dfs(root *TreeNode, cur, pre *int) {
	if root == nil {
		return
	}
	dfs(root.Left, cur, pre)
	if *pre != -1 {
		*cur = min(*cur, root.Val-*pre)
	}
	*pre = root.Val
	dfs(root.Right, cur, pre)

}
