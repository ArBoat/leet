func singleNumber(nums []int) int {
    res:=0
    for i:=0;i<32;i++ {
        value:=0
        bit:=1<<i
        for _, v:= range nums {
            // &
            if v&bit != 0 {
                value++
            }
        }
        // IM: |=
        if value%3 != 0 {
            res |= bit
        }
    }
    return res
}

/*
每位与运算 %3
*/