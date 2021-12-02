func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	for l < r {
		m := l + (r-l)/2
		if check(matrix, k, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func check(matrix [][]int, k int, mid int) bool {
	n := len(matrix)
	i, j := n-1, 0
	count := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count >= k
}