package problem0086

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lesserHead, greaterHead := &ListNode{}, &ListNode{}
	lesserCurrent, greaterCurrent := lesserHead, greaterHead

	for head != nil {
		if head.Val < x {
			lesserCurrent.Next = head
			lesserCurrent = lesserCurrent.Next
		} else {
			greaterCurrent.Next = head
			greaterCurrent = greaterCurrent.Next
		}
		head = head.Next
	}

	greaterCurrent.Next = nil
	lesserCurrent.Next = greaterHead.Next

	return lesserHead.Next
}
