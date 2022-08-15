func countSubstrings(s string) int {
	res := 0
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if (j-i <= 1) || (dp[i+1][j-1]) {
					dp[i][j] = true
					res++
				}
			}
		}
	}
	return res
}

/*

dp[i][j] :  bool
dp[i][j] = false (s[i] != s[j])
         = true (s[i] == s[j]&& j-i<=1 ) res++
         = true (s[i] == s[j]&& j-i>1&& dp[i+1][j-1] ) res++
dp[i][i] = false
*/