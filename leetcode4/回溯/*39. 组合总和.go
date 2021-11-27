func combinationSum(candidates []int, target int) [][]int {
    sort.Slice(candidates, func(a, b int)bool {
       return candidates[a]<candidates[b]
    })
    res:=[][]int{}
    path:=[]int{}
    var backtrack func(path []int, target int, start int)
    backtrack = func(path []int, target int, start int) {
        if target == 0 {
            //IM: go copy 闭包
            temp:=make([]int, len(path))
            copy(temp, path)
            // temp
            res = append(res, temp)
            return
        }
        for i:=start;i<len(candidates);i++{
            // IM: 剪枝
            if candidates[i]<=target {
                // IM: candidates[i]
                backtrack(append(path, candidates[i]),target-candidates[i],i)
            } 
        }
    }
    backtrack(path, target, 0)
    return res
}