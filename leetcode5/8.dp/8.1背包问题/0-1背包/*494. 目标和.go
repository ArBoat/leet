func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	//IM: abs(sum) < target, target may negtive
	if (target+sum)%2 == 1 || sum < abs(target) {
		return 0
	}
	p := (target + sum) / 2

	dp := make([]int, p+1)
	//IM : init dp[0] = 1
	dp[0] = 1
	for _, num := range nums {
		for j := p; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[p]
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

/*
positive - negetive = target
positive + negetive = sum
positive = (target + sum)/2 // (target + sum) is  even

dp[j]:  count of set which sum is positive
dp[j] = dp[j] + dp[j-nums[i]]
dp[0] = 1
*/