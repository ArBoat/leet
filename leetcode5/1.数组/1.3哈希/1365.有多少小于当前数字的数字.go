func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	rawNums := make([]int, n)
	copy(rawNums, nums)
	// 内置排序
	sort.Ints(nums)
	m := make(map[int]int)
	// IM: 从后向前遍历
	for i := n - 1; i >= 0; i-- {
		m[nums[i]] = i
	}
	for i, v := range rawNums {
		res[i] = m[v]
	}
	return res
}