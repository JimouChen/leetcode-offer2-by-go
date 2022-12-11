package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	first, sec := head, head
	for i := 0; i < k; i++ {
		sec = sec.Next
	}
	for sec != nil {
		first = first.Next
		sec = sec.Next
	}
	return first
}
