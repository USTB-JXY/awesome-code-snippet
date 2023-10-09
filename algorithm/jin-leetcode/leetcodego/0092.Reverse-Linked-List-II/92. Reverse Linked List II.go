package leetcode

import (
	"fmt"

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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}

func myreverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m > n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre == nil {
		return newHead.Next
	}
	// 【input】:{[1 2 3 4 5] 1 5}

	// 【guocheng】:[2 1 3 4 5]
	//
	// 【guocheng】:[3 2 1 4 5]
	//     h    c
	// 【guocheng】:[4 3 2 1 5]
	// 【guocheng】:[5 4 3 2 1]
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		temp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = temp
		fmt.Printf("【guocheng--】:%v \n", structures.List2Ints(newHead.Next))
	}
	return newHead.Next
}
