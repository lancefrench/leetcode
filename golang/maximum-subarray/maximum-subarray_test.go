package problem0053

import (
	"testing"
)

type tc struct {
	params
	output
}

type params struct {
	nums []int
}

type output struct {
	ans int
}

func TestProblem0053(t *testing.T) {
	tcs := []tc{
		{
			params{
				[]int{},
			},
			output{},
		},
	}
}
