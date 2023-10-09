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

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}

func mylowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := mylowestCommonAncestor236(root.Left, p, q)
	right := mylowestCommonAncestor236(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
