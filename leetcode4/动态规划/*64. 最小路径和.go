func minPathSum(grid [][]int) int {
    if len(grid)<1 || len(grid[0])<1 {return 0}
    m, n:= len(grid), len(grid[0])
    //IM: [][]int
    dp:=make([][]int, m)
    for i:=range dp {
        dp[i] = make([]int, n)
    }
    // do not forget ,  location
    dp[0][0] = grid[0][0]
    for i:=1;i<m;i++{
        dp[i][0] = dp[i-1][0]+grid[i][0]
    }
    for j:=1;j<n;j++{
        dp[0][j] = dp[0][j-1]+grid[0][j]
    }
    for i:=1;i<m;i++{
        for j:=1;j<n;j++{
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }
    return dp[m-1][n-1]
}

func min(a, b int) int {
    if a>b{
        return b
    }
    return a
}

/*
dp[i][j] 求 dp[m-1][n-1]

dp[0][0] = grid[0][0]
dp[i][0] = dp[i-1][0]+grid[i][0]
dp[0][j] = dp[0][j-1]+grid[0][j]

dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
*/