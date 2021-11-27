func rob(nums []int) int {
    n:=len(nums)
    dp:=make([]int, n+1)
    dp[0]=0
    dp[1]=nums[0]
    for i:=2;i<=n;i++ {
        dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
    }
    return dp[n]
}

func max(a, b int) int {
    if a>b{
        return a
    }
    return b
}

/*
dp[i] 在0-i范围内最大金额
d[0] = 0
d[1] = nums[0]

dp[i] = max(dp[i-1],nums[i]+dp[i-2])
dp[n]
*/