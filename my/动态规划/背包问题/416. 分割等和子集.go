func canPartition(nums []int) bool {
    // 重要
    if len(nums)<2 { return false}
    sum:=0
    for _,v:=range nums {
        sum+=v
    }
    if sum%2 == 1 { return false}
    sum/=2
    dp:=make([]bool, sum+1)
    dp[0]=true
    for _,v:=range nums{
        //会更新dp[0]
        for i:= sum; i>=0; i-- {
            if i>=v{
                dp[i] = dp[i] || dp[i-v]
            }
        }
    }
    return dp[sum]
}

/*
0-1, nums， target， 反
dp[i], 目标和为i是为可取到
dp[0]=true // 要不要更新
dp[i] = dp[i] || dp[i-v]
*/