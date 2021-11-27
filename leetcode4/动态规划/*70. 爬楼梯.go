func climbStairs(n int) int {
    if n==0 || n == 1 { 
        return 1
    }
    p,q,r:=1,1,2
    for i:=2;i<=n;i++{
        r = p + q
        p = q
        q = r
    }
    return r
}