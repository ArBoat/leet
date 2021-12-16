func getSum(a int, b int) int {
    xor:= a ^ b
    and:= (a & b) << 1
    for and != 0 {
        temp:= xor & and
        xor = xor ^ and
        and = temp << 1
    }
    return xor
}

/*
位运算
无进位加法使用异或运算计算得出
进位结果使用与运算和移位运算计算得出
循环此过程，直到进位为 0
*/

xor:= a ^ b
and:= (a & b) << 1