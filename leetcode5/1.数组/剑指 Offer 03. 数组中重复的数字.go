func findRepeatNumber(nums []int) int {
	for i := range nums {
		if nums[i] > len(nums)-1 || nums[i] < 0 {
			return -1
		} // 数字都在 0～n-1 的范围内, 沟通需不需要
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

/*
slice遍历改变值用index

1. i = v: 跳过
2. v = num[v]: 返回
3. 否则， 交换v和nums[v]
返回-1

遍历每个位置：
1. 如果发现当前元素num已经躺在正确的位置上，continue;
2. 否则，即num不在正确位置上：
  1）如果发现其正确的位置上已经躺着num了，即发现重复元素，return num;
  2）否则，开启循环移位流程，将num发往正确位置上，其怼出来的元素，继续发往它应该去的地方，如此循环。。。

*/