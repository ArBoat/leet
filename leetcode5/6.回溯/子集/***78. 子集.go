func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	res := [][]int{}
	path := []int{}
	var backtrack func(path []int, cur int)
	backtrack = func(path []int, cur int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := cur; i < n; i++ {
			backtrack(append(path, nums[i]), i+1)
		}
	}
	backtrack(path, 0)
	return res
}