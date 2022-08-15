func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	res := [][]int{}
	path := []int{}
	var backtrack func(path []int, index int)
	backtrack = func(path []int, index int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			backtrack(append(path, nums[i]), i+1)
		}
	}
	backtrack(path, 0)
	return res
}

/*
1. 先排序
2. 判断重复
*  测试用例 加顺序
*/
