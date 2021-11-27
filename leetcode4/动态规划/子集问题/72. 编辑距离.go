func minDistance(word1 string, word2 string) int {
    m,n := len(word1), len(word2)
    dp:= make([][]int, m+1)
    for i:= range dp {
        dp[i] = make([]int, n+1)
    }
    dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
                dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
            }
        }
    }
    return dp[m][n]
}

func min(a, b int) int {
    if a>b {
        return b
    }
    return a
}

/*
m, n
dp[i][j] 表示两个单词的遍历的指针， 求dp[m-1][n-1]
dp[0][0] = 0
dp[0][j] = j // j>0
dp[i][0] = i // i>0

dp[i][j] = dp[i-1][j-1]  // 相等
    min (
    dp[i-1][j-1]+1   // 替换
    dp[i-1][j]+1    // 删除
    dp[i][j-1] +1    // 插入
)
*/