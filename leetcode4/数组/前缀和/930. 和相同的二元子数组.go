func numSubarraysWithSum(nums []int, goal int) int {
    res:=0
    hash:=make(map[int]int)
    hash[0]=1
    preSum:=0
    for _, v:= range nums {
        preSum+=v
        if hash[preSum-goal]>0 {
            res+=hash[preSum-goal]
        }
        hash[preSum]++
    }
    return res
}

prefixSum[j]−prefixSum[i−1]==k