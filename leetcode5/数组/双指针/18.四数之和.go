func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	l := len(nums)
	if l < 4 {
		return res
	}
	sort.Ints(nums)
	for m := 0; m < l-3; m++ {
		if m > 0 && nums[m] == nums[m-1] {
			continue
		}
		for n := m + 1; n < l-2; n++ {
			if n > m+1 && nums[n] == nums[n-1] {
				continue
			}
			i, j := n+1, l-1
			for i < j {
				sum := nums[m] + nums[n] + nums[i] + nums[j]
				if sum < target {
					i++
					for i < j && nums[i] == nums[i-1] {
						i++
					}
				} else if sum > target {
					j--
					for i < j && nums[j] == nums[j+1] {
						j--
					}
				} else {
					res = append(res, []int{nums[m], nums[n], nums[i], nums[j]})
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
	}
	return res
}

// 三数之和进阶