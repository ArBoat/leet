// l<=r
func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }
    // IM: x/2处即可确定平方根
    l, r := 1, x/2
    for l<=r{
        m:=l+(r-l)/2
        if m*m==x{
            return m
        } else if m*m>x{
            r = m-1
        } else {
            l = m+1
        }
    }
    // IM: return r
    return r
}