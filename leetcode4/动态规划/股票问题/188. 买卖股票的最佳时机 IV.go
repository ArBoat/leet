func maxProfit(k int, prices []int) int {
    n:=len(prices)
    if n<1 { return 0 }
    if k>n/2{
        k=n/2
    }
    dp:=make([][]int, 2)
    for i:=range dp {
        dp[i] = make([]int, k+1)
    }
    // IM:初始化
    for i:=0;i<=k;i++{
        dp[1][i] = math.MinInt32
    }
    for _,v:=range prices {
        for i:=1;i<=k;i++{
            dp[0][i] = max(dp[0][i], dp[1][i]+v)
            dp[1][i] = max(dp[1][i], dp[0][i-1]-v)
        }
    }
    return dp[0][k]
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*
dp[i][k][0] = max(dp[i][k][0], dp[i-1][k][1]+v)
dp[i][k][1] = max(dp[i][k][1], dp[i-1][k-1][0]-v)

dp[0][0]=0
dp[1][0] = min

*/