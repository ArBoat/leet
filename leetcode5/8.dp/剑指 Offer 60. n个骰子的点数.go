func dicesProbability(n int) []float64 {
	dp := make([]float64, 6)
	for i := range dp {
		dp[i] = float64(1) / float64(6) //IM : float64
	}
	for i := 2; i <= n; i++ {
		tmp := make([]float64, 5*i+1)
		for j := range dp {
			for k := 0; k < 6; k++ { // 0-5
				tmp[j+k] += dp[j] / 6
			}
		}
		dp = tmp
	}
	return dp
}

/*
5*i+1
*/