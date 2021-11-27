func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, m*n)
	// 赋值正确
	l, r, t, b := 0, n-1, 0, m-1
	for index := 0; index < m*n; {
		// IM : 循环中每次都要判断index<m*n
		//l to r
		for i := l; i <= r && index < m*n; i++ {
			res[index] = matrix[t][i]
			index++
		}
		t++
		// t to b
		for i := t; i <= b && index < m*n; i++ {
			res[index] = matrix[i][r]
			index++
		}
		r--
		// r to l
		for i := r; i >= l && index < m*n; i-- {
			res[index] = matrix[b][i]
			index++
		}
		b--
		// b to t
		for i := b; i >= t && index < m*n; i-- {
			res[index] = matrix[i][l]
			index++
		}
		l++
	}
	return res
}