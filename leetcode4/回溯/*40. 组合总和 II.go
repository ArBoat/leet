func combinationSum2(candidates []int, target int) [][]int {
    sort.Slice(candidates, func(a, b int)bool {
       return candidates[a]<candidates[b]
    })

    // IM: []int{}
    res:=[][]int{}
    path:=[]int{}
    var backtrack func(path []int, target int, start int)
    backtrack = func(path []int, target int, start int) {
        if target == 0 {
            temp:=make([]int, len(path))
            copy(temp, path)
            res = append(res, temp)
            return
        }
        for i:=start;i<len(candidates);i++{
            // IM : 去重
            if (i > start && candidates[i] == candidates[i - 1]) {
                continue
            }
            // <=
            if candidates[i]<=target {
                backtrack(append(path, candidates[i]),target-candidates[i],i+1)
            } 
        }
    }
    backtrack(path, target, 0)
    return res
}

// 