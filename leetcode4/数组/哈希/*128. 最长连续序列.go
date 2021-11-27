func longestConsecutive(nums []int) int {
    hash:=make(map[int]bool)
    for _, v:= range nums {
        hash[v] = true
    }
    res:=0
    for i:= range hash {
        if !hash[i-1] {
            cur:=i
            curLen:=1
            for hash[cur+1] {
                cur++
                curLen++
            }
            if res < curLen {
                res = curLen
            }
        }
    }
    return res
}
/*
O(n)/O(n)
哈希表
*/