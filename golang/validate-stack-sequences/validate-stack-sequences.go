package problem0946

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	pushedIndex := 0
	for _, v := range popped {
		for len(stack) == 0 || stack[len(stack)-1] != v {
			if pushedIndex == len(pushed) {
				return false
			}
			stack = append(stack, pushed[pushedIndex])
			pushedIndex++
		}
		stack = stack[:len(stack)-1]
	}
	return true
}
