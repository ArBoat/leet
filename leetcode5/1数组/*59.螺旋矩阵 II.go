func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	for num := 1; num <= n*n; {
		//l to r
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		// t to b
		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--
		// r to l
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--
		// b to t
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}