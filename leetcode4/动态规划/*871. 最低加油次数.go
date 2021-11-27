func minRefuelStops(target int, startFuel int, stations [][]int) int {
    n:=len(stations)
    dp:=make([]int, n+1)
    dp[0] = startFuel
    for i:=0;i<n;i++{
        for t:=i;t>=0;t--{
            if dp[t] >= stations[i][0] {
                // 如果选择了
                dp[t+1] = max(dp[t+1], dp[t]+stations[i][1])
            }
        }
    }
    for i:=range dp {
        if dp[i] >= target {
            return i
        }
    }
    return -1
}

func max(a, b int) int {
    if a>b{
        return a
    }
    return b
    
}

/*
dp[i]为i次油最远距离， 求dp[i]>=target的最小i

*/