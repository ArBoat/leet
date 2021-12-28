func reverseStr(s string, k int) string {
	sbyte := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k { // 循环条件
		if i+k < n {
			reverse(sbyte[i : i+k])
		} else {
			reverse(sbyte[i:])
		}
	}
	return string(sbyte)
}

func reverse(a []byte) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}