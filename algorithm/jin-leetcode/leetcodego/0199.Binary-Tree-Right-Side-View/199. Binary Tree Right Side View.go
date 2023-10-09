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

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val)
		queue = queue[n:]
	}
	return res
}
func myrightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	quene := []*TreeNode{root}
	n := len(quene)
	for i := 0; i < n; i++ {
		if quene[i].Left != nil {
			quene = append(quene, quene[i].Left)
		}
		if quene[i].Right != nil {
			quene = append(quene, quene[i].Right)
		}
		if i == n-1 {
			n = len(quene)
			res = append(res, quene[i].Val)
		}
	}
	return res

}
