func subarraySum(nums []int, k int) int {
    res:=0
    hash:=make(map[int]int)
    hash[0]=1
    preSum:=0
    for _, v:= range nums {
        preSum+=v
        res += hash[preSum-k]
        hash[preSum]++
    }
    return res
}

/*
O(n)/O(n)
0:1
preSum

p(x) = n[0]+..n[x]
n[x] = p(x) - p(x-1)
p(-1) = 0
i...j: k
p(j) - p(i-1) = k




*/