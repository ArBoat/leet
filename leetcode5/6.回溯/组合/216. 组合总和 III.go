func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(path []int, index int)
	backtrack = func(path []int, index int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			// IM:add check
			sum := 0
			for _, v := range temp {
				sum += v
			}
			if sum == n {
				res = append(res, temp)

			}
			// IM: return
			return
		}
		//IM: 9-i+1>= k-len(path)
		for i := index; i <= 9+len(path)-k+1; i++ {
			backtrack(append(path, i), i+1)
		}
	}
	backtrack(path, 1)
	return res
}