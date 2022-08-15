func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	stack := []int{} // save value
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[nums2[i]] = stack[len(stack)-1]
		} else {
			mp[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = mp[v]
	}
	return res
}

/*
单调栈（递减）+哈希表
“亲们记住，一但要求下一个更大的元素，就是用单调栈解，力扣题库相似的题目都是这个解法。”
for for for for
*/