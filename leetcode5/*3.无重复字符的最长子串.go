func lengthOfLongestSubstring(s string) int {
	win := make(map[byte]int)
	left := 0
	res := 0
	for right := range s {
		win[s[right]]++
		for win[s[right]] > 1 {
			win[s[left]]--
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