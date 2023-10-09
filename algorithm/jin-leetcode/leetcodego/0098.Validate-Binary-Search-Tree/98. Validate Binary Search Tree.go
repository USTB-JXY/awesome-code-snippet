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

// 解法一，直接按照定义比较大小，比 root 节点小的都在左边，比 root 节点大的都在右边
func isValidBST(root *TreeNode) bool {
	return isValidbst(root, math.Inf(-1), math.Inf(1))
}
func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}

// 解法二，把 BST 按照左中右的顺序输出到数组中，如果是 BST，则数组中的数字是从小到大有序的，如果出现逆序就不是 BST
func isValidBST1(root *TreeNode) bool {
	arr := []int{}
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}

func myisValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if root.Right != nil {
			return (root.Val > root.Left.Val) && (root.Val < root.Right.Val) && myisValidBST(root.Left) && myisValidBST(root.Right)
		}
		return root.Val > root.Left.Val && myisValidBST(root.Left)
	}
	if root.Right != nil {
		return root.Val < root.Right.Val && myisValidBST(root.Right)
	}
	return true
}
func my2isValidBST(root *TreeNode) bool {
	return ifdfs(root, math.Inf(1), math.Inf(-1))
}
func ifdfs(root *TreeNode, maxv, minv float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)

	return v < maxv && v > minv && ifdfs(root.Left, v, minv) && ifdfs(root.Right, maxv, v)
}
