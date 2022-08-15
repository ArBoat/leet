func wordBreak1(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, v := range wordDict {
			if i >= len(v) && v == s[i-len(v):i] {
				dp[i] = dp[i] || dp[i-len(v)]
			}
		}
	}
	return dp[len(s)]
}

/*
有序完全背包
dp[j]
dp[i] = dp[i] || d[i-len(v)]
dp[0] = true : 为了推导
*/

func wordBreak2(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, v := range wordDict {
		wordSet[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

/*
有序
dp[j]
d[i] = dp[j]&&wordSet[j,i]
dp[0] = true : 为了推导
但本题还有特殊性，因为是要求子串，最好是遍历背包放在外循环，将遍历物品放在内循环。
*/