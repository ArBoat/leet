func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	res := 1
	preDiff, curDiff := 0, 0
	for i := 0; i < n-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (curDiff > 0 && preDiff <= 0) || (curDiff < 0 && preDiff >= 0) {
			preDiff = curDiff
			res++
		}
	}
	return res
}

/*
确认n<2情况
cur 不判断0
pre 判断0(初始情况)

时间复杂度：$O(n)$
空间复杂度：$O(1)$
*/