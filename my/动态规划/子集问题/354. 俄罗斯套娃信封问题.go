func maxEnvelopes(envelopes [][]int) int {
    if len(envelopes) == 0 {return 0}
    // 第一个顺序，相等第二个逆序
    sort.Slice(envelopes, func(a, b int)bool{
       return envelopes[a][0] < envelopes[b][0] || 
       (envelopes[a][0]==envelopes[b][0] && envelopes[a][1]>envelopes[b][1])
    })
    n:=len(envelopes)
    dp:=make([]int, n)
    res:=0
    for i:=range dp {
        dp[i] = 1
    }
    for i:=0;i<n;i++{
        for j:=0;j<i;j++{
            if envelopes[i][1] > envelopes[j][1] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        res = max(res, dp[i])
    }
    return res
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*
sort, 然后求最长子递增序列
*/