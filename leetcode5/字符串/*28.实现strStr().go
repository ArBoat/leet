func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	//next 的 ptm(Partial Match Table)
	next := make([]int, m)
	next[0] = -1                  // IMPORTANT: 初始化
	for i, j := 0, -1; i < m-1; { // IM: j = -1, i< m-1
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	// 匹配
	i, j := 0, 0
	for i < n && j < m {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		return i - j
	}
	return -1
}

//因为 KMP 利用已匹配部分中相同的「前缀」和「后缀」来加速下一次的匹配
//因为 KMP 的原串指针不会进行回溯（没有朴素匹配中回到下一个「发起点」的过程）
//当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//https://www.zhihu.com/question/21923021/answer/281346746
