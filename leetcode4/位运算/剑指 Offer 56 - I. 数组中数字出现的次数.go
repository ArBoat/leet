func singleNumbers(nums []int) []int {
    //全部异或
    xor:=0
    for _, v:= range nums {
        xor ^= v
    }
    // 找到第一个1
    // IM: mask & xor
    mask:=1
    for mask & xor == 0 {
        mask<<=1
    }
    // 分组
    x, y := 0,0
    for _, v:= range nums {
        if v & mask == 0 {
            x^=v
        } else {
            y^=v
        }
    }
    return []int{x, y}
}