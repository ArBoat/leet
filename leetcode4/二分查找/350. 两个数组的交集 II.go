func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    i, j := 0, 0
    res:= []int{}
    for i<len(nums1) && j < len(nums2) {
        // IM: 小的先走
        if nums1[i] > nums2[j]{
            j ++ 
        } else if nums1[i] < nums2[j] {
            i ++ 
        } else {
            //先加入res
            res = append(res, nums1[i])
            i++
            j++
        }
    }
    return res
}

/*
O(mlogm+nlogn)/ O(min(m,n))
*/