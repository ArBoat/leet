//数学分析
func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	return res * n
}

//dp
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ { //IM: j 是段  j <= i/2
			dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[2] = 1
dp[i] 长度为n的绳子的最大乘积

dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j))) // i-j>1 //j是段
*/