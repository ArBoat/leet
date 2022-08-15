func findMaxForm(strs []string, m int, n int) int {
	//m ,n 取值
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zero, one := 0, 0
		for _, v := range []byte(str) {
			if v == '0' {
				zero++ // 确认只有1，0
			} else {
				one++
			}
		}
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
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
dp[i][j]: 最多有i个0，j个1的子集最大长度
dp[i][j] = max(dp[i][j], dp[i - zeroNum][j - oneNum] + 1)

*/