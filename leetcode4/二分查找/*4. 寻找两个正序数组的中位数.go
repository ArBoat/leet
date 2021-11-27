func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l:=len(nums1) + len(nums2)
    if l%2 == 0 {
        return float64(findKth(nums1, nums2, l/2)+
        findKth(nums1, nums2, l/2 +1))/2.0
    } else {
        return float64(findKth(nums1, nums2, l/2+1))
    }
    // l/2+1
}
// index ， k在变
func findKth(nums1 , nums2 []int, k int) int {
    index1, index2:= 0, 0
    for {
        // IM: 终止三条件
        if index1 == len(nums1) {
            // index2 + k -1
            return nums2[index2 + k -1]
        }
        if index2 == len(nums2) {
            return nums1[index1 + k - 1]
        }
        // ***
        if k == 1 {
            return min(nums1[index1], nums2[index2])
        }
        // IM: new ***
        new1:= min(len(nums1), index1 + k/2) -1
        new2:= min(len(nums2), index2 + k/2) -1
        
        //index2 + k -1
        if nums1[new1]<= nums2[new2] {
            // ***
            k = index1+k-1-new1
            index1 = new1 + 1
        } else {
            k = index2+k-1-new2
            index2 = new2 + 1
        }
        // 更新 index 和 k
    }
}

func min(a, b int) int {
    if a<b {
        return a
    }
    return b
}