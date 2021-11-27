func findPeakElement(nums []int) int {
//    if len(nums) == 0 { return -1}  IM: 没必要
    l, r := 0, len(nums)-1
    for l < r {
        mid := l + (r-l)/2
        if nums[mid] > nums[mid+1] {
            r = mid 
        } else {
            l = mid +1
        }
    }
    return l
}

/*
O(logn)/O(1)
*/