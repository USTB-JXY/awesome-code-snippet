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

func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	inorder230(root, k, &count, &res)
	return res
}

func inorder230(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		inorder230(node.Left, k, count, ans)
		*count++
		if *count == k {
			*ans = node.Val
			return
		}
		inorder230(node.Right, k, count, ans)
	}
}

func mykthSmallest(root *TreeNode, k int) int {
	res, time := 0, 0
	myorder(root, &res, &time, k)
	//fmt.Println(k, time, res)
	if time == k {
		return res
	}
	return -1

}
func myorder(root *TreeNode, res, time *int, k int) {
	if root == nil {
		return
	}
	myorder(root.Left, res, time, k)
	if *time == k {
		return
	}
	*time++
	if *time == k {
		*res = root.Val
		return
	} else {
		myorder(root.Right, res, time, k)
	}

}
