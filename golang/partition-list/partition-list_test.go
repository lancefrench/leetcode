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
	head *ListNode
	x    int
}

type output struct {
	ans *ListNode
}

func slice2linked(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}
	return res
}

func Test_Problem0086(t *testing.T) {
	tcs := []tc{
		{
			params{},
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
