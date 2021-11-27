func maxSubArray(nums []int) int {
    n:=len(nums)
    if n == 0 { return 0}
    res:=nums[0]
    dp:=make([]int, n)
    //IM: 初始化
    dp[0] = nums[0]
    //IM: i:=1,  nums[i], nums[i]+dp[i-1]
    for i:=1;i<n;i++ {
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        res = max(res, dp[i])
    }
    // return res
    return res
}



func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*
前序和
*/

func maxSubArray(nums []int) int {
    n:=len(nums)
    if n == 0 {
        return 0
    }
    pre, res := nums[0], nums[0]
    for i:=1;i<n;i++ {
        pre = max(nums[i], nums[i]+pre)
        res = max(res, pre)
    }
    return res
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}