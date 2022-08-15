func movingCount(m int, n int, k int) int {
	res := 0
	board := make([][]bool, m)
	for i := range board {
		board[i] = make([]bool, n)
	}
	board[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if digitSum(i)+digitSum(j) <= k && canReach(board, m, n, i, j) {
				board[i][j] = true
				res++
			}
		}
	}
	fmt.Println(board)
	return res
}

func digitSum(number int) int {
	sum := 0
	for number > 0 { //im :number>0
		sum += number % 10
		number /= 10
	}
	return sum
}

func canReach(board [][]bool, m, n, i, j int) bool {
	return board[i][j] ||
		(j-1 >= 0 && board[i][j-1]) ||
		(i-1 >= 0 && board[i-1][j]) ||
		(j+1 < n && board[i][j+1]) ||
		(i+1 < m && board[i+1][j])
}