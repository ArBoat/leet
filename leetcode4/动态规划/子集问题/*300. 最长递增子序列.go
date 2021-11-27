// 单调队列+二分法
func lengthOfLIS(nums []int) int {
    if len(nums) == 0 || len(nums) == 1 {
        return len(nums)
    }
    sub:=[]int{}
    for _,v := range nums {
        // IM: len(sub) == 0
        if len(sub) == 0 || v>sub[len(sub)-1] {
            sub = append(sub, v)
        } else {
            l, r:= 0, len(sub)
            for l<r {
                m:=l+(r-l)/2
                // IM: sub[m], not nums[m]
                if sub[m] >= v{
                     r=m
                } else {
                    l = m+1
                }
            }
            sub[l] = v
        }
    }
    return len(sub)
}

//dp

func lengthOfLIS(nums []int) int {
    res:=0
    n:=len(nums)
    dp:=make([]int, n)
    for i:= range dp {
        dp[i] = 1
    }
    for i:=0; i< n;i++ {
        for j:=0;j<i;j++{
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        res= max(res, dp[i])
    }
    return res
}
func max(a, b int) int {
    if a>b{
        return a
    }
    return b
}