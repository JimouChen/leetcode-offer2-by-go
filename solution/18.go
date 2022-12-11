package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	pre, cur := newHead, newHead.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}
	return newHead.Next
}
