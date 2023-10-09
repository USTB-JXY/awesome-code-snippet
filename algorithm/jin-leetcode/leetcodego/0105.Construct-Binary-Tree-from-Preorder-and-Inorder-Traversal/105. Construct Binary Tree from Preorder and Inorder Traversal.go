package leetcode

import (
	"fmt"

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
//  For example, given

//  preorder = [3,9,20,15,7]
//  inorder = [9,3,15,20,7]

// Return the following binary tree:

//		 3
//		/ \
//	   9  20
//		 /  \
//		15   7
//
// 解法一, 直接传入需要的 slice 范围作为输入, 可以避免申请对应 inorder 索引的内存, 内存使用(leetcode test case) 4.7MB -> 4.3MB.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(preorder[1:pos+1], inorder[:pos])
			root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
			break
		}
	}
	return root
}
func mybuildTree(preorder []int, inorder []int) *TreeNode {
	fmt.Printf("【preorder】:%v   【inorder】:%v     \n", preorder, inorder)
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for pos, val := range inorder {
		if val == root.Val {
			root.Left = mybuildTree(preorder[1:pos+1], inorder[:pos])
			root.Right = mybuildTree(preorder[pos+1:], inorder[pos+1:])
			break
		}
	}
	return root
}

// 解法二
func buildTree1(preorder []int, inorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildPreIn2TreeDFS(preorder, 0, len(preorder)-1, 0, inPos)
}

func buildPreIn2TreeDFS(pre []int, preStart int, preEnd int, inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: pre[preStart]}
	rootIdx := inPos[pre[preStart]]
	leftLen := rootIdx - inStart
	root.Left = buildPreIn2TreeDFS(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = buildPreIn2TreeDFS(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
	return root
}
