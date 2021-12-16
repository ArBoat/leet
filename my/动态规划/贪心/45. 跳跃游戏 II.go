func jump(nums []int) int {
    distance, count, end:= 0, 0, 0
    //IM: 边界位i<len(nums)-1，而不是整个数组
    for i:=0;i<len(nums)-1;i++{
        distance = max(distance, i+nums[i])
        if end == i {
            count++
            end = distance
        }
    }
    return count
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

/*
用end记录每次最大起跳位置
*/