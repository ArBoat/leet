
func subarraysDivByK(nums []int, k int) int {
    res:=0
    hash:=make(map[int]int)
    hash[0]=1
    sum:=0
    for _, v:=range nums{
        sum+=v
        mod:= (sum%k+k)%k
        if hash[mod]>0{
            res+=hash[mod]
        }
        hash[mod]++
    }
    return res
}

//O(n)/O(n) go 负数取余数为负，要处理下

//O(n)/O(n) go 负数取余数为负，要处理下