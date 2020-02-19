package problem0019

/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode 是节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre = head
	var pos = head
	for i := 0; i < n-1; i++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head.Next
	}
	pre = pre.Next
	for pre.Next != nil {
		pos = pos.Next
		pre = pre.Next
	}
	pos.Next = pos.Next.Next
	return head
}
