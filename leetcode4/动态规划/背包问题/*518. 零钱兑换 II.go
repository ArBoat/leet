func change(amount int, coins []int) int {
    dp:=make([]int, amount+1)
    dp[0] = 1
    for _,v:=range coins {
        for i:=1;i<=amount;i++ {
            if i>=v {
                dp[i] = dp[i] + dp[i-v]
            }
        }
    }
    return dp[amount]
}

/*

dp[i] 表示总金额为i的组合数
dp[0] = 1 //不会更新
dp[i] = dp[i] + dp[i-v]

完全无序， nums,target,正
*/