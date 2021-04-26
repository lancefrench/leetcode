package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	params
	output
}

type params struct {
	l1 []int
	l2 []int
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

func TestProblem0002(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{2, 4, 3},
				[]int{5, 6, 4},
			},
			output{
				[]int{7, 0, 8},
			},
		},
		{
			params{
				[]int{0},
				[]int{0},
			},
			output{
				[]int{0},
			},
		},
		{
			params{
				[]int{9, 9, 9, 9, 9, 9, 9},
				[]int{9, 9, 9, 9},
			},
			output{
				[]int{8, 9, 9, 9, 0, 0, 0, 1},
			},
		},
	}

	for _, tc := range tcs {
		p, o := tc.params, tc.output

		// Because it's easier to initialize our test
		// cases as slices of ints, we use a helper
		// function to load them into a ListNode for
		// testing.
		inputList1 := sliceToListNode(p.l1)
		inputList2 := sliceToListNode(p.l2)
		actual := addTwoNumbers(inputList1, inputList2)
		expected := sliceToListNode(o.ans)

		assert.Equal(t, expected, actual)
	}
}
