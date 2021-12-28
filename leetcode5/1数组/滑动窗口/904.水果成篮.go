func totalFruit(fruits []int) int {
	win := make(map[int]int)
	res := 0
	left := 0
	for right, v := range fruits {
		win[v]++
		for len(win) > 2 {
			win[fruits[left]]--
			if win[fruits[left]] == 0 {
				delete(win, fruits[left])
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}