func translateNum(num int) int {
    s:=strconv.Itoa(num)
    n:=len(s)
    if n==0 || n==1 { return 1}
    dp:=make([]int, n+1)
    // init dp value
    dp[0] = 1
    dp[1] = 1
    for i:=2;i<=n;i++{
        dp[i] = dp[i-1]
        //IM: condition
        if s[i-2] != '0' && 10*(s[i-2]-'0')+(s[i-1]-'0')<=25 {
            dp[i]+=dp[i-2]
        }
    }
    return dp[n]
}

/*
n:=0,1 1种
i:=1
dp[i] = dp[i-1]
dp[i] += dp[i-2]  10<=num[i-2]*10+nums[i-1]<=25 
*/