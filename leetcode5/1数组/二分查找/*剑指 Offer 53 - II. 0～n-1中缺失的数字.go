func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == m {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

/*
O(logN)/
二分法查找 “右子数组的首位元素”
	排序数组中的搜索问题，首先想到 二分法 解决
*/