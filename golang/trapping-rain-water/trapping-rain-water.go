package problem0042

type stack []int

func (s *stack) Push(x int) {
	*s = append(*s, x)
}

func (s *stack) Pop() (int, bool) {
	l := len(*s)
	if l == 0 {
		return 0, false
	}

	i := l - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

func (s *stack) Peek() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func trap(height []int) int {
	var stack stack
	var res int

	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack.Peek()] {
			top, _ := stack.Pop()
			if len(stack) == 0 {
				break
			}
			distance := i - stack.Peek() - 1
			boundedHeight := min(height[i], height[stack.Peek()]) - height[top]
			res += distance * boundedHeight
		}
		stack.Push(i)
	}
	return res
}
