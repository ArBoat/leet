func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)
	reverse(b, 0, len(b)-1)
	return string(b)
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

/*
1.反转区间为前n的子串
2.反转区间为n到末尾的子串
3.反转整个字符串
*/