func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	// IM:=1
	count := 1
	end := points[0][1]
	for _, v := range points {
		//IM : >end
		if v[0] > end {
			count++
			end = v[1]
		}
	}
	return count
}

/*
看区间可不可重叠
*/