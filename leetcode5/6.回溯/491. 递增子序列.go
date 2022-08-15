func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(path []int, index int)
	backtrack = func(path []int, index int) {
		if len(path) > 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		hash := make(map[int]bool)
		for i := index; i < len(nums); i++ {
			//当前i比path最后一个大
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || hash[nums[i]] {
				continue
			}
			hash[nums[i]] = true
			backtrack(append(path, nums[i]), i+1)
		}
	}
	backtrack(path, 0)
	return res
}