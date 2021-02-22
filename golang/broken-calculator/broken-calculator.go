package Problem0991

func brokenCalc(X int, Y int) int {
	res := 0
	for X < Y {
		res += Y%2 + 1
		Y = (Y + 1) / 2
	}
	return res + X - Y
}
