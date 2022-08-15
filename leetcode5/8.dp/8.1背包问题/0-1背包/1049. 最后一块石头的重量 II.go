func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, stone := range stones {
		for j := target; j >= stone; j-- {
			dp[j] = max(dp[j], dp[j-stone]+stone)
		}
	}
	return sum - dp[target] - dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[j]容量为j的背包最多可以装的石头重量
dp[j] = max(dp[j], dp[j-stones[i]]+ stones[i])
dp[0] = 0
target = sum/2
sum - dp[target] -dp[target]
*/