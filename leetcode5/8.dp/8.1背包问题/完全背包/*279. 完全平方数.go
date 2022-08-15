func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 0; i*i <= n; i++ {
		for j := i * i; j <= n; j++ { //IM: j:=i*i
			dp[j] = min(dp[j], dp[j-i*i]+1) //IM: j:=i*i
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
无序完全背包
dp[j]
dp[j] = min(dp[j], dp[j-nums[i]]+1)
dp[j] = 0, others MaxInt32

*/