func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1 // init
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}

/*
dp[i][j] : s, t s在自己上寻找t
dp[i][j] = dp[i-1][j]  (s[i-1] != s[j-1])
         = dp[i-1][j] + dp[i-1][j-1]  [no,yes]    (s[i-1] == s[j-1])
dp[i][0] = 1
dp[0][j] = 0
*/     