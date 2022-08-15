func maxProfit(prices []int) int {
	dp0, dp1 := 0, math.MinInt32
	pre := 0
	for _, v := range prices {
		temp := dp0
		dp0 = max(dp0, dp1+v)
		dp1 = max(dp1, pre-v)
		pre = temp
	}
	return dp0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])   still， sell
dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])   still， buy


pre = 0 , dp0 = 0, dp1 =math.MinInt32
temp = dp0
dp0 =max(dp0, dp1+v)
dp = max(dp0, pre - v)
pre = temp

*/