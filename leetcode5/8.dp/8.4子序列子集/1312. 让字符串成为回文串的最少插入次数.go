func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- { // im: i--
		for j := i + 1; j < n; j++ { // j++
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
dp[0][len(s)-1]

中间向两边, i--, j++

s[i] == s[j] :  dp[i][j] = dp[i+1][j-1]
!=  : dp[i][j] = min(dp[i+1][j],dp[i][j-1]) +1
*/