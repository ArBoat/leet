func eraseOverlapIntervals(intervals [][]int) int {
	//check zero len
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1 // 非交叉区间数量
	end := intervals[0][1]
	for _, v := range intervals {
		// im: >=
		if v[0] >= end {
			count++
			end = v[1]
		}
	}
	return len(intervals) - count
}

/*
重叠子区间， 边界根据题意
*/