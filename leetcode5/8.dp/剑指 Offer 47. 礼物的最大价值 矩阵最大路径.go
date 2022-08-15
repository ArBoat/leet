func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = grid[i-1][j-1] + max(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
 dp[i][j] = grid[i-1][j-1] + max(dp[i-1][j], dp[i][j-1])
*/