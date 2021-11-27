func maxProduct(nums []int) int {
    if len(nums) == 0 { return 0}
    maxCur, minCur, res:= nums[0],nums[0],nums[0]
    // IM: 从1开始
    for i:=1;i<len(nums);i++{
        maxTemp, minTemp:=maxCur, minCur
        // nums[i], nums[i]*max, nums[i]*min
        maxCur = max(nums[i], max(nums[i]*maxTemp, nums[i]*minTemp))
        minCur = min(nums[i], min(nums[i]*maxTemp, nums[i]*minTemp))
        res = max(res, maxCur)
    }
    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}