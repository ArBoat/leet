func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi, si := 0, 0
	for ; gi < len(g) && si < len(s); si++ {
		if s[si] >= g[gi] {
			gi++
		}
	}
	return gi
}

// 思路1：优先考虑胃饼干
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)