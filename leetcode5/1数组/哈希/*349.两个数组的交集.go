func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]struct{})
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := hash[v]; !ok {
			hash[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := hash[v]; ok {
			res = append(res, v)
			delete(hash, v)
		}
	}
	return res
}

//hash两次遍历