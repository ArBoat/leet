func numWays(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    p, q, r:= 1, 1, 2
    for i:=2;i<=n;i++{
        r = (p+q)%1000000007
        p = q
        q = r
    } 
    return r
}
// 斐波那契数列
// 此题要求取1000000007的模