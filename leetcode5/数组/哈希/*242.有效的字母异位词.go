func isAnagram(s string, t string) bool {
	// IM：注意判断长度相等
	if len(s) != len(t) {
		return false
	}
	hash := make(map[rune]int)
	for _, v := range s {
		hash[v]++
	}
	for _, v := range t {
		hash[v]--
		if hash[v] < 0 {
			return false
		}
	}
	return true
}