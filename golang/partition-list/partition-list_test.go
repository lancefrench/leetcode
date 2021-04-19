package problem0086

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			output{
				[]int{1, 2, 2, 4, 3, 5},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		// Because it's easier to initialize our test
		// cases as slices of ints, we use a helper
		// function to load them into a ListNode for
		// testing.
		inputList := sliceToListNode(p.head)
		got := partition(inputList, p.x)
		want := sliceToListNode(o.ans)

		assert.Equal(t, want, got)
	}
}
