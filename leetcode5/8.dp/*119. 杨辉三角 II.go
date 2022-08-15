func getRow(rowIndex int) []int {
	dp := make([]int, rowIndex+1)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i <= rowIndex; i++ { // i  <=
		for j := i - 1; j >= 1; j-- { // j >= 反向防覆盖遍历
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp
}

/*
dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
dp[i][0] = 1
dp[i][i] = 1

1
11
121
1331
*/