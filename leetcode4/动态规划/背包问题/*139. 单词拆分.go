func wordBreak(s string, wordDict []string) bool {
    n:=len(s)
    dp:=make([]bool, n+1)
    dp[0] = true
    for i:=1;i<=n;i++ {
        for _, v:= range wordDict {
            if  i>=len(v) &&  s[i-len(v):i] == v {
                dp[i] = dp[i] || dp[i-len(v)]
                // abcs, [c, bc]
            }
        }
    }
    return dp[n]
}

/*
完全有序背包   先target，后选择，正
dp[0] = true
dp[i] = dp[i] || dp[i-len(v)]
*/