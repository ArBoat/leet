func isPerfectSquare(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	l, r := 1, num/2
	for l <= r {
		m := l + (r-l)/2
		if m*m == num {
			return true
		} else if m*m < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}