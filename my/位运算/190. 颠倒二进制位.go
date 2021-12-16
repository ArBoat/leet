func reverseBits(num uint32) uint32 {
    res:=uint32(0)
    for i:=32; i>0; i-- {
        res = res<<1
        res += num & 1
        num = num>>1
    }
    return res
}

/*
位运算
结果左移动， num右移， 加上num最后是否是一的一位
*/