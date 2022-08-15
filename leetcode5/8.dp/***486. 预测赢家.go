func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
   dp[i][j]表示从i到j范围 玩家1 能拿到的净积分，赢过对方积分
   i>j dp[i][j]=0
   dp[i][i] = nums[i]
   i<j : dp[i][j]=max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])// 减去对手的积分
*/