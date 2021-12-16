func missingNumber(nums []int) int {
    res:= len(nums)
    for i,v:=range nums{
        res ^= i^v
    }
    return res
}

/*
O(1)
位运算
*/