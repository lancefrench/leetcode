package problem0012

type mapping struct {
	num int
	str string
}

var roman []mapping = []mapping{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	i, res := 0, ""
	for num > 0 {
		if num-roman[i].num >= 0 {
			res += roman[i].str
			num -= roman[i].num
		} else {
			i++
		}
	}
	return res
}
