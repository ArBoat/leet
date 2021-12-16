func longestPalindromeSubseq(s string) int {
    n:=len(s)
    dp:=make([][]int, n)
    for i:=range dp{
        dp[i] = make([]int, n)
        //IM : dp[i][i] = 1
        dp[i][i] = 1
    }
    // IM: i:=n-2;i>=0;i--
    for i:=n-2;i>=0;i--{
        for j:=i+1;j<n;j++ {
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            }
        }
    }
    return dp[0][n-1]
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*
dp[i][j]
求dp[0][n-1]
dp[i][i]=1
i--, j++
dp[i][j] = dp[i+1][j-1] + 2 // s[i] == s[j]
         =  max(dp[i+1][j],dp[i][j-1])

*/