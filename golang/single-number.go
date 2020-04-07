func singleNumber(nums []int) int {
	var unique int = nums[0]
	for _, v := range nums[1:] {
		unique ^=  v
	}
	return unique
}
