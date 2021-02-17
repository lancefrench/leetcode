package problem0011

import (
	"testing"
)

type params struct {
	one []int
}

type ans struct {
	one int
}

func Test_Problem0011(t *testing.T) {
	input := params{[]int{1, 2, 3, 1}}
	got := maxArea(input.one)
	want := 3

	if got != want {
		t.Errorf("maxArea() = %d; want %d", got, want)
	}
}
