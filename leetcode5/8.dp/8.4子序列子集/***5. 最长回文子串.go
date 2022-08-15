func longestPalindrome(s string) string {
	res := ""
	for i := range s {
		s1 := find(s, i, i)
		s2 := find(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func find(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r] // im: [l+1:r]
}

/*
从中心向两边扩展， 分为奇数和偶数两种情况，少数比dp更优解
*/