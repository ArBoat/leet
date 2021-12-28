//hash
func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	res := make([]int, 0)
	for _, v := range nums1 {
		hash[v]++
	}
	for _, v := range nums2 {
		if count := hash[v]; count > 0 {
			res = append(res, v)
			hash[v]--
		}
	}
	return res
}

// 双指针
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	res := []int{}
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}