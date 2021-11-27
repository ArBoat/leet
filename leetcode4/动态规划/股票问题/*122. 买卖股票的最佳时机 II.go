func maxProfit(prices []int) int {
    dp0, dp1:=0, math.MinInt32
    for _, v:= range prices {
        temp:= dp0
        dp0 = max(dp0, dp1+v)
        dp1 = max(dp1, temp-v)
    }
    return dp0
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*
k无穷

dp[i][0] = max(dp[i][0], dp[i-1][1]+v)
dp[i][1] = max(dp[i][1], dp[i-1][0]-v)
dp[0][0] = 0
dp[0][1] = min
*/