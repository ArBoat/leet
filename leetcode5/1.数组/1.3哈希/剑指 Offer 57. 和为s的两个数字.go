//一.哈希
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{target - v, v}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

// 二.双指针
func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] == target {
			return []int{nums[l], nums[r]}
		} else if nums[l]+nums[r] < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}

/*
一.哈希
类似1.两数之和， 但是这里返回值为元素值而不是索引

二.双指针 类似二分法： 排序数组

*/