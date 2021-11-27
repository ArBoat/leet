func singleNumber(nums []int) int {
    res:=0
    for _, v:= range nums{
        res ^= v
    }
    return res
}

/*
位运算 
异或
a ^ a = 0
a ^ 0 = a
*/