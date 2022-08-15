//DP + 贪心 + 二分
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		if len(dp) == 0 || dp[len(dp)-1] < num {
			dp = append(dp, num)
		} else {
			l, r := 0, len(dp)
			for l < r {
				m := l + (r-l)/2
				if dp[m] >= num {
					r = m
				} else {
					l = m + 1
				}
			}
			dp[l] = num
		} //二分查找,特殊,  [,)
	}
	fmt.Println(dp)
	return len(dp)
}

/*
查找第一个大于等于num的index
*/

//dp
func lengthOfLIS(nums []int) int {
	res := 0
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1 //IM: Init
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ { //IM: i<j
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
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
dp[i]
dp[i] = max(dp[i], dp[j]+1)  // num[i]>num[j]， j<=i
init: dp[i] = 1

res= max(res, dp[i])
*/