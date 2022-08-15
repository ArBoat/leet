func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(a, b int) bool {
		return candidates[a] < candidates[b]
	})
	res := [][]int{}
	path := []int{}
	var backtrack func(path []int, target int, start int)
	backtrack = func(path []int, target int, start int) {
		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			v := candidates[i]
			if v <= target {
				// IM: 可重复
				backtrack(append(path, v), target-v, i)
			}
		}
	}
	backtrack(path, target, 0)
	return res
}

/*
O(S), 所有解长度和
O(target)
*/