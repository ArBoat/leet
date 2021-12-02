func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	sumLeft := 0
	// 前缀和
	for i, v := range nums {
		// 要用*而不是/, /不是整数
		if sumLeft*2+nums[i] == sum {
			return i
		}
		sumLeft += v
	}
	return -1
}

/*
sl + nums[i] + sr = s
sl = sr
sl = (s-nums[i])/2
sl * 2 + nums[i] = s
*/ 
