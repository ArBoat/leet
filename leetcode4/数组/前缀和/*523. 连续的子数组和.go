func checkSubarraySum(nums []int, k int) bool {
    sum:=0
    hash:=make(map[int]int)
    // 存索引,存最小
    hash[0]=-1
    for i, v:=range nums{
        sum+=v
        var mod int
        if k == 0 {
            mod = sum
        } else {
            mod = sum%k
        }
        if index, ok:= hash[mod];ok{
            if i - index > 1 {
                return true
            }
        } else {
            hash[mod] = i
        }
    }
    return false
}
/*
O(n)/O(n)
1. 存最小的下标，
2.hash[0]=1
*/