func combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(path []int, cur int)
	backtrack = func(path []int, cur int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// im: prune
		for i := cur; i <= n-k+len(path)+1; i++ {
			backtrack(append(path, i), i+1) // i+1
		}
	}
	backtrack(path, 1)
	return res
}

/*
n-i+1 >= k-len(path)
i<= n-k+len(path)+1
*/

/*
O((n k)*k)
O(n)
*/