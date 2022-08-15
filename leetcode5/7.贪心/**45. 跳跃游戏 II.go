func jump(nums []int) int {
	distance, count, end := 0, 0, 0
	//end为当前最远
	//IM: 边界位i<len(nums)-1，而不是整个数组
	for i := 0; i < len(nums)-1; i++ {
		distance = max(distance, i+nums[i])
		if i == end {
			count++
			end = distance
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
一定能跳到最远

用end记录每次最大起跳位置
*/