package problem0206

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode
	var pos *ListNode
	if head == nil {
		return nil
	}
	cur = head
	pos = cur.Next
	for pos != nil {
		cur.Next = pre
		pre = cur
		cur = pos
		pos = pos.Next
	}
	cur.Next = pre
	return cur
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}
