func maxProfit(prices []int) int {
	dp0, dp1 := 0, math.MinInt32
	for _, price := range prices {
		dp0 = max(dp0, dp1+price)
		dp1 = max(dp1, -price)
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
通用公式， k=1
dp[i][0] =  max(dp[i-1][0], dp[i-1][1]+prices[i])
dp[i][1] =  max(dp[i-1][1], -prices[i])
dp[0][0] = 0
dp[0][1] = MinInt32

dp0 =max(dp0, dp1 + price[i])
dp1 =max(dp1, -price[i])
dp0 = 0
dp1 = MinInt32

*/