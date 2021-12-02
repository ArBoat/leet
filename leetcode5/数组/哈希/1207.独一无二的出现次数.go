func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	// struct{} 占位当set
	n := make(map[int]struct{})
	for _, v := range m {
		if _, ok := n[v]; ok {
			return false
		}
		n[v] = struct{}{}
	}
	return true
}