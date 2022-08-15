//l<=r
func search(nums []int, target int) int {
    //IM: 边界
    if len(nums) == 0 { return -1 }
    l, r:= 0, len(nums)-1
    for l<=r {
        mid := l + (r-l)/2
        // IM: 三个分支
        if nums[mid] == target {
            return mid
        } else if nums[mid] >= nums[0]{
            if nums[mid] > target && target >= nums[0]{
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if nums[mid] < target && target < nums[0]{
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    // IM: 找不到
    return -1
}
/*
O(logn)/O(1)
二分查找
注意边界

往两边逼
*/