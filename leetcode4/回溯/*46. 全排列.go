func permute(nums []int) [][]int {
    res:=[][]int{}
    path:=[]int{}
    // IM: used
    used:=make(map[int]bool)
    var backtrack func(path []int, used map[int]bool)
    backtrack = func(path []int, used map[int]bool) {
        if len(path) == len(nums) {
            tmp:=make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
        }
        for i,v:= range nums {
            if !used[i]{
                used[i] = true
                backtrack(append(path, v), used)
                used[i] = false
            }
        }
    }
    backtrack(path, used)
    return res
}