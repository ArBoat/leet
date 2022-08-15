func repeatedSubstringPattern(s string) bool {
	n := len(s)
	next := make([]int, n+1)
	next[0] = -1 // 当前字符串位置之前最大重复前缀后一位
	for i, j := 0, -1; i < n; {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	if next[n] > 0 && n%(n-(next[n])) == 0 {
		return true
	} else {
		return false
	}
}

// next 多存最后一个位置 当前字符串位置之前最大重复前缀长度， 以及后一位位置
// next[n]>0 && n (n(next[n])) == 0  