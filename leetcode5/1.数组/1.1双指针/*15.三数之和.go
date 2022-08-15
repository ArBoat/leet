func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	n := len(nums)
	for k := 0; k < n-2; k++ {
		i, j := k+1, n-1
		if nums[k] > 0 { //直接break
			break
		}
		if k > 0 && nums[k] == nums[k-1] { // k=0 是第一个， 肯定和前面不同，也没有前面
			continue
		}
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if sum < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				} // 判断<j很重要
			} else if sum > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}

	}
	return res
}