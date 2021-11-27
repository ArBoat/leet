func fib(n int) int {
    if n == 0 || n == 1 { 
        return n
    }
	p, q, r:= 0, 1, 1
    for i:=2;i<=n;i++{
        r = (p+q)%1000000007
        p = q
        q = r
    }
    return r
}

// 斐波那契数列
// 此题要求取1000000007的模