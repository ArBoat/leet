func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	used := make(map[int]bool)
	sort.Ints(nums)
	var backtrack func(path []int, used map[int]bool)
	backtrack = func(path []int, used map[int]bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i, v := range nums {

			if !used[i] {
				// 三个条件
				if i > 0 && nums[i] == nums[i-1] && used[i-1] {
					continue
				}
				used[i] = true
				backtrack(append(path, v), used)
				used[i] = false
			}
		}
	}
	backtrack(path, used)
	return res
}