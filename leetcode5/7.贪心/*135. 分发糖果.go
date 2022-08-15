func candy(ratings []int) int {
	//init
	n := len(ratings)
	candyList := make([]int, n)
	for i := range candyList {
		candyList[i] = 1
	}
	// left to right
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candyList[i] = candyList[i-1] + 1
		}
	}
	// right to left
	for j := n - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			candyList[j] = max(candyList[j], candyList[j+1]+1)
		}
	}
	sum := 0
	for _, v := range candyList {
		sum += v
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
遍历两次
1. 从左到右：
2. 从右到左：
*/