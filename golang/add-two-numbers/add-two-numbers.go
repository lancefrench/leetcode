package problem0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{Val: 0, Next: nil}
	carry := 0
	current := &res

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		sum := x + y + carry
		current.Next = &ListNode{Val: sum % 10, Next: nil}
		carry = sum / 10

		// Move forward one node in the lists
		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		if carry != 0 {
			current.Next = &ListNode{Val: 1, Next: nil}
		}

	}
	return res.Next
}
