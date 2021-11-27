func combinationSum4(nums []int, target int) int {
    dp:=make([]int, target+1)
    dp[0]=1
    for i:=1;i<=target;i++ {
        for _,v:=range nums {
            if i>=v{
                dp[i] = dp[i] + dp[i-v]
            }
        }
    }
    return dp[target]
}



/*
完全，有序， target,nums,,正
 dp[i] 表示目标为i时的组合数
 dp[0]=1
 dp[i] = dp[i] + dp[i-v]
*/