package problem0011

import (
	"testing"
)

type example struct {
	params
	answer
}

type params struct {
	one []int
}

type answer struct {
	one int
}

func Test_Problem0011(t *testing.T) {
	examples := []example{
		{
			params{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			answer{49},
		},
		{
			params{[]int{1, 1}},
			answer{1},
		},
		{
			params{[]int{4, 3, 2, 1, 4}},
			answer{16},
		},
		{
			params{[]int{1, 2, 1}},
			answer{2},
		},
	}

	for _, ex := range examples {
		a, p := ex.answer, ex.params

		got := maxArea(p.one)
		want := a.one

		if got != want {
			t.Errorf("got maxArea(%v) = %d; want %d", ex.params, got, want)
		}
	}
}
