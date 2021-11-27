func cuttingRope(n int) int {
    dp:=make([]int, n+1)
    for i:=2;i<=n;i++{
        maxCur:=0
        for j:=1;j<=i/2;j++{
            maxCur = max(maxCur, max(j*(i-j), j*dp[i-j]))
        }
        dp[i] = maxCur
    }
    return dp[n]
}

func max(a, b int) int {
    if a>b{
        return a
    }
    return b
}