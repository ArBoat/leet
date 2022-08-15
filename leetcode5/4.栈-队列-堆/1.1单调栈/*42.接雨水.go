func trap(height []int) int {
	l, r := 0, len(height)-1 // index
	lMax, rMax := 0, 0       // value
	res := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > lMax {
				lMax = height[l]
			} else {
				res += lMax - height[l]
			}
			// IM : l++
			l++
		} else {
			if height[r] > rMax {
				rMax = height[r]
			} else {
				res += rMax - height[r]
			}
			// IM: r--
			r--
		}
	}
	return res
}

/*
每一列雨水的高度，取决于该列左侧最高的柱子和右侧最高的柱子中最矮的那个柱子的高度
*/