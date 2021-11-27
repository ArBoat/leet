func findTargetSumWays(nums []int, target int) int {
    sum:=0
    for _,v:=range nums {
        sum+=v
    }
    // IM: sum<target
    if (sum+target)%2==1 || sum<target { return 0 }
    //
    t:=(sum+target)/2
    dp:=make([]int, t+1)
    dp[0] = 1
    for _,v:=range nums {
        for i:=t;i>=0;i--{
            if i>=v{
                dp[i] = dp[i] + dp[i-v]
            }
        }
    }
    return dp[t]
}

/*
0-1背包   选择， target， 反
A-B=target
A+B=nums
A = (target+nums)/2, (target+nums)一定为偶数
dp[i] 目标为i的个数
dp[0]=1
dp[i] = dp[i] + dp[i-v]



0-1, nums, target,反
*/