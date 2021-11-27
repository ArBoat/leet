func divide(dividend int, divisor int) int {
    if dividend == (-1<<31) && divisor == -1 {
        return (1<<31) - 1
    }
    res:=0
    dividendAbs:= abs(dividend)
    divisorAbs:= abs(divisor)
    for dividendAbs >= divisorAbs {
        power:=0
        for dividendAbs - divisorAbs << power >= 0 {
            power++
        }
        res += 1 << (power - 1)
        dividendAbs -= divisorAbs << (power - 1)
    }
    //注意符号
    if (dividend > 0) == (divisor >= 0){
        return res
    }
    return -res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

/*
O(logn)
17/3
*/