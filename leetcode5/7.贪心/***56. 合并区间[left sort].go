func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(a, b int) bool {
		// 按左边排
		if intervals[a][0] == intervals[b][0] {
			return intervals[a][1] < intervals[b][1]
		}
		return intervals[a][0] < intervals[b][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
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
按左边排
1.[[1,4],[4,5]]
2.[[3,4],[2,4]]
3.[[1,4],[2,3]]
4.[[2,3],[4,5],[6,7],[8,9],[1,10]]
*/