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

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func myhasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

//    // using a map to store the nodes we walked
//    bool hasCycle01(ListNode *head) {
// 	unordered_map<ListNode*,bool > m;
// 	ListNode* p = head;
// 	m[(ListNode*)NULL] = true;
// 	while( m.find(p) == m.end() ) {
// 		m[p] = true;
// 		p = p -> next;
// 	}
// 	return p != NULL;
// }

// // Change the node's of value, mark the footprint by a special value
// bool hasCycle02(ListNode *head) {
// 	ListNode* p = head;
// 	// using INT_MAX as the mark could be a bug!
// 	while( p &&  p->val != INT_MAX ) {
// 		p->val = INT_MAX;
// 		p = p -> next;
// 	}
// 	return p != NULL;
// }

// /*
//  * if there is a cycle in the list, then we can use two pointers travers the list.
//  * one pointer traverse one step each time, another one traverse two steps each time.
//  * so, those two pointers meet together, that means there must be a cycle inside the list.
//  */
// bool hasCycle03(ListNode *head) {
// 	if (head==NULL || head->next==NULL) return false;
// 	ListNode* fast=head;
// 	ListNode* slow=head;
// 	do{
// 		slow = slow->next;
// 		fast = fast->next->next;
// 	}while(fast != NULL && fast->next != NULL && fast != slow);
// 	return fast == slow? true : false;
// }

// // broken all of nodes
// bool hasCycle04(ListNode *head) {
// 	 ListNode dummy (0);
// 	ListNode* p = NULL;

// 	while (head != NULL) {
// 		p = head;
// 		head = head->next;

// 		// Meet the old Node that next pointed to dummy
// 		//This is cycle of linked list
// 		if (p->next == &dummy) return true;

// 		p->next = &dummy; // next point to dummy
// 	}
// 	return false;
// }
