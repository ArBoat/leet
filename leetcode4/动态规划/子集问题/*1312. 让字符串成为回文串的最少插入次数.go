func minInsertions(s string) int {
    n:=len(s)
    // 能不能取到空
    dp:=make([][]int, n)
    for i:=range dp {
        dp[i] = make([]int, n)
    }
    for i:=n-1;i>=0;i--{
        for j:=i+1;j<n;j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1]
            } else {
                dp[i][j] = min(dp[i+1][j],dp[i][j-1]) + 1
            }
        }
    }
    return dp[0][n-1]
}

func min(a, b int) int {
    if a>b {
        return b
    }
    return a
}

/*
dp[i][j], 在i，j范围内插入多少
i>=j; dp[i][j] = 0

s[i] == s[j] :  dp[i][j] = dp[i+1][j-1]
!=  : dp[i][j] = min(dp[i+1][j],dp[i][j-1]) +1
*/