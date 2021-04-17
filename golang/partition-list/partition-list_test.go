package Problem0086

import "testing"

type tc struct {
	params
	output
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type params struct {
	head []int
	x    int
}

type output struct {
	ans []int
}

func (n *ListNode) append(val int) {
	newNode := ListNode{val, nil}

	for n.Next != nil {
		n = n.Next
	}
	n.Next = &newNode
}

func sliceToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	linkedList := &ListNode{
		Val: nums[0],
	}

	for i := 1; i < len(nums); i++ {
		linkedList.append(nums[i])
	}
	return linkedList
}

func TestProblem0086(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{1, 4, 3, 2, 5, 2},
				3,
			},
			output{},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		got := partition(p.head, p.x)
		want := o.ans

		if got != want {
			t.Errorf("got partition(\"%v\") = %d want %d", p.head, got, want)
		}
	}
}
