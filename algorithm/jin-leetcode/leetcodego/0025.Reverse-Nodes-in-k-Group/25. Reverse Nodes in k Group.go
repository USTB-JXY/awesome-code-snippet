package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 1 2 3 nil
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := myodreverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func myreverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := myodreverse(head, node)
	head.Next = myreverseKGroup(node, k)
	return newHead
}

// p tf       p   f t   p      f t
// 1 2 3 4       1 2 3 4    1 4 2 3   2 1 4  3 - 3 2 1 4
// 2 1 3 4
// 3 2 1 4
// 4 3 2 1
func odreverse(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}
func myodreverse(first *ListNode, last *ListNode) *ListNode {
	//从[head ,last) -1翻转元素
	prev, cur := last, first
	for cur != last {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

// ```go
// //双指针
// func reverseList(head *ListNode) *ListNode {
//     var pre *ListNode
//     cur := head
//     for cur != nil {
//         next := cur.Next
//         cur.Next = pre
//         pre = cur
//         cur = next
//     }
//     return pre
// }

// //递归
// func reverseList(head *ListNode) *ListNode {
//     return help(nil, head)
// }

// func help(pre, head *ListNode)*ListNode{
//     if head == nil {
//         return pre
//     }
//     next := head.Next
//     head.Next = pre
//     return help(head, next)
// }
