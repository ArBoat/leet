func rotate(nums []int, k int) {
	// IM : 注意取余
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

/*
F1:额外数组复制赋值: O(n)/O(n)
F2:多次翻转：O(n)/O(1)
左转后转整个数组， 右转先翻转整个数组

左： 反转区间为前n的子串
    反转区间为n到末尾的子串
    反转整个字符串

右： 反转整个字符串
    反转区间为前k的子串
    反转区间为k到末尾的子串
*/