func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	//IM: %, /
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[j]: dp[j]表示 背包总容量是j，最大可以凑成j的子集总和为dp[j]
dp[j] = max(dp[j], d[j-nums[i]] + nums[i])
dp[0] = 0
*/