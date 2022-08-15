func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		//IM : dp[i][i] = 1
		dp[i][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2 // IM: i+1;j-1;+2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[i][j] : dp[0][n-1]

dp[i][j] = dp[i+1][j-1]+2  (s[i] == s[j])        //IM : +2
         = max(dp[i+1][j], dp[i][j-1])
i--, j++
dp[i][i] = 1
*/