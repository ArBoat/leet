func maxProfit(prices []int) int {
	dp20, dp21 := 0, math.MinInt32
	dp10, dp11 := 0, math.MinInt32
	for _, v := range prices {
		dp20 = max(dp20, dp21+v)
		dp21 = max(dp21, dp10-v)
		dp10 = max(dp10, dp11+v)
		dp11 = max(dp11, -v)
	}
	return dp20
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+v)
dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-v)
dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+v)
dp[i][1][1] = max(dp[i-1][1][1], -v)

dp20, dp21 = 0, min
dp10, dp11 = 0, min

dp20 = max(dp20, dp21+v)
dp21 = max(dp21, dp10-v)
dp10 = max(dp10, dp11+v)
dp11 = max(dp11, -v)
*/