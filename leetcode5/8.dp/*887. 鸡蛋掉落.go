func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1) // 一层一层最坏
	}
	m := 0
	for dp[k][m] < n { //im: <n
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i-1][m-1] + dp[i][m-1] + 1
		}
	}
	return m
}

/*
dp[k][m] k个鸡蛋m次能测试的高度, 求m
dp[0][..] =0
dp[..][0] = 0
dp[k][m] = 碎了(下楼)+没碎（上楼）+1
         = dp[k-1][m-1] + dp[k][m-1] +1


*/