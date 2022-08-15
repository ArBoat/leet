func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32 //IM: init
	}
	dp[0] = 0
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
无序
dp[j] 凑成总金额为i的最少硬币数
dp[j] = min(dp[j], dp[j-coin[i]]+1)
dp[0] = 0, others maxInt32
*/