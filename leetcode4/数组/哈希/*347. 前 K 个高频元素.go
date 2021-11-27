func topKFrequent(nums []int, k int) []int {
    hash:=make(map[int]int)
    res:=[]int{}
    for _,v :=range nums{
        hash[v]++
    }
    for k:= range hash {
        res = append(res, k)
    }
    sort.Slice(res, func(i, j int) bool {
		return hash[res[i]] > hash[res[j]]
	})
    return res[:k]
}