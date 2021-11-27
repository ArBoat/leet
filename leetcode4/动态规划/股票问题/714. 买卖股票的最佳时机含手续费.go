func maxProfit(prices []int, fee int) int {
    if len(prices)<1 { return 0}
    dp0, dp1:= 0, math.MinInt32
    for _, v:= range prices{
        temp := dp0
        dp0 = max(dp0, dp1+v)
        dp1 = max(dp1, temp-v-fee)
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
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)

*/