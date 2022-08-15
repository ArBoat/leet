func partitionLabels(s string) []int {
	res := []int{} // 每个字符串片段的长度
	record := [26]int{}
	// record last index
	for i, v := range s {
		record[v-'a'] = i
	}
	left, right := 0, 0
	for i, v := range s {
		right = max(right, record[v-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
统计每一个字符最后出现的位置
从头遍历字符，并更新字符的最远出现下标，如果找到字符最远出现位置下标和当前下标相等了，则找到了分割点
*/