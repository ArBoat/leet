
func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		//IM: still != '0'
		if i >= 2 && s[i-2] != '0' && 10*(s[i-2]-'0')+(s[i-1]-'0') <= 25 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

/*
dp[i] =dp[i-1] + dp[i-2]
*/