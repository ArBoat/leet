func permuteUnique(nums []int) [][]int {
    res:=[][]int{}
    path:=[]int{}
    used:=make(map[int]bool)
    sort.Ints(nums)
    var backtrack func(path []int, used map[int]bool)
    backtrack = func(path []int, used map[int]bool) {
        if len(path) == len(nums) {
            tmp:=make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
        }
        for i := 0; i < len(nums); i++ {
            // IM: 去重条件, ****
            if i-1>=0 && nums[i] == nums[i-1] && used[i-1]{
                continue
            }
            if !used[i]{
                used[i] = true
                backtrack(append(path, nums[i]), used)
                used[i] = false
            }
        }
    }
    backtrack(path, used)
    return res
}