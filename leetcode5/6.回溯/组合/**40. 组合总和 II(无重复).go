func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(a, b int) bool {
		return candidates[a] < candidates[b]
	})

	res := [][]int{}
	path := []int{}
	var backtrack func(path []int, target, start int)
	backtrack = func(path []int, target, start int) {
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			v := candidates[i]
			//IM: prune1
			if i > start && v == candidates[i-1] {
				continue
			}
			//IM: prune2
			if v <= target {
				backtrack(append(path, v), target-v, i+1) //IM: i+1
			}
		}
	}
	backtrack(path, target, 0)
	return res
}

/*
集合（数组candidates）有重复元素，但还不能有重复的组合
*/