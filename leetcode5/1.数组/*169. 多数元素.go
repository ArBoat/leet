func majorityElement(nums []int) int {
	count, target := 0, 0
	for _, v := range nums {
		if count == 0 {
			target = v
		}
		if target == v {
			count++
		} else {
			count--
		}
	}
	return target
}

/*
1.hash统计
2.排序中间数
3.摩尔投票

一定大于 ⌊ n/2 ⌋ 的元素， 等于不行
Boyer-Moore 投票算法
*/