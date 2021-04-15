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

func Test_Problem0086(t *testing.T) {
	tcs := []tc{
		{
			params{},
			output{},
		},
	}
}
