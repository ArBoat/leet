func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	//最后一步不算花费
	return min(dp[len(dp)-1], dp[len(dp)-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
dp[i] = min(dp[i - 1], dp[i - 2]) + cost[i]
dp[i]的定义：到达第i个台阶所花费的最少体力为dp[i]
*/