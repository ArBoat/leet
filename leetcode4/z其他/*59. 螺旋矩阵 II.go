func generateMatrix(n int) [][]int {
    res:=make([][]int,n)
    for i:=range res {
        res[i] = make([]int, n)
    }
    l, r, t, b := 0, n-1, 0, n-1
    num:=1
    for num<=n*n {
        // l to r
        for i:=l;i<=r;i++ {
            res[t][i] = num
            num++
        }
        t++
        // t to b
        for i:=t;i<=b;i++ {
            res[i][r] = num
            num++
        }
        r--
        // r to l
        for i:=r;i>=l;i--{
            res[b][i] = num
            num++
        }
        b--
        // b to t
        for i:=b;i>=t;i--{
            res[i][l] = num
            num++
        }
        l++
    }
    return res
}