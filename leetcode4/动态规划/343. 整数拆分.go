func integerBreak(n int) int {
    dp:= make([]int, n+1)
    for i:=2;i<=n;i++{
        maxCur:=0
		// j<=i/2
        for j:=1;j<=i/2;j++{
            maxCur = max(maxCur,max(j*(i-j), j*dp[i-j]))
        }
        dp[i] = maxCur
    }
    return dp[n]
}

func max(a, b int) int {
    if a>b {
        return a
    } 
    return b
}

/*
dp[n]表示将n拆分✖成连个正整数的成集最大
dp[0] = dp[1] = 0
maxCur = max(maxCur,max(j*(i-j),j*dp[i-j]))
             不拆了    继续拆
dp[i] = maxCur

求dp[n]
*/