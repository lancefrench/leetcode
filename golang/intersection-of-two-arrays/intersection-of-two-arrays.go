package problem0349

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool, len(nums1))
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		set[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		_, found := set[nums2[i]]
		if found {
			res = append(res, nums2[i])
			delete(set, nums2[i])
		}
	}
	return res
}
