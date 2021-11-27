func subsets(nums []int) [][]int {
    if len(nums) == 0 { return nil}
    res:=[][]int{}
    path:=[]int{}
    var backtrack func(path []int, cur int)
    backtrack = func(path []int, cur int){
        // go 参数加入res要copy
        tmp:= make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)

        for i := cur; i<len(nums);i++ {
            backtrack(append(path, nums[i]), i+1)
        }
    }
    backtrack(path,0)
    return res
}

/*
O(n*2^n)/O(n)
回溯
*/