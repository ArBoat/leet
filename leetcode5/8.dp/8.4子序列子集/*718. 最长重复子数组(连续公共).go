func findLength(nums1 []int, nums2 []int) int {
	res := 0
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] { //IM:  i-1, j-1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[i][j]
dp[i][j] = dp[i-1][j-1]+1 // nums1[i-1] = nums2[j-1]
         = 0
init 0
res = max(res, dp[i][j])
*/