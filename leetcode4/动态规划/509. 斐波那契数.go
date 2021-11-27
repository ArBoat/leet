func fib(n int) int {
    if n<2 { return n}
    p, q, r := 0,1,1
    for i:=2;i<=n;i++{
        r = p + q
        p = q
        q = r
    }
    return r
}
