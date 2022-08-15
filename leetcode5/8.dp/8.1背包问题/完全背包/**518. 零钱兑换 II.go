func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coin]
		}
	}
	return dp[amount]
}

/*
无序
dp[j]总金额为i的组合数 dp[amount]
dp[j] = dp[j] + dp[j-coins[i]
dp[0]=1
*/