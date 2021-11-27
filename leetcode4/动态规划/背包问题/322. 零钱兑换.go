func coinChange(coins []int, amount int) int {
    dp:=make([]int, amount+1)
    for i:=1;i<=amount;i++ {
        dp[i] = math.MaxInt32
    }
    for _,v :=range coins {
        for i:=1;i<=amount;i++ {
            if i>=v {
                dp[i] = min(dp[i], dp[i-v]+1)
            }
        }
    }
    // IM: 返回-1
    if dp[amount] == math.MaxInt32 {
        return -1
    }
    return dp[amount]
}

func min(a, b int) int {
    if a>b {
        return b
    }
    return a
}


/*
完全背包， 无序， 硬币最小数量
先选择，后目标，正
dp[i]表示总金额为amount所需最小硬币数量
dp[0]=0, dp[1-amount] == max
dp[i] = min(dp[i], dp[i-v]+1)

判断最后-1
*/