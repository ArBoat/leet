// 多指针
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	i, j := 0, n-1
	for p := j; p >= 0; p-- {
		if m, n := nums[i]*nums[i], nums[j]*nums[j]; m > n {
			res[p] = m
			i++
		} else {
			res[p] = n
			j--
		}
	}
	return res
}