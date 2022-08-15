func monotoneIncreasingDigits(n int) int {
	if n == 0 {
		return 0
	}
	s := []byte(strconv.Itoa(n))
	// im: i>0
	for i := len(s) - 1; i > 0; i-- {
		if s[i-1] > s[i] {
			s[i-1] -= 1
			for j := i; j < len(s); j++ { //im: j<len(s)
				s[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(s))
	return res
}

/*
本题只要想清楚个例，例如98，一旦出现strNum[i - 1] > strNum[i]的情况（非单调递增），首先想让strNum[i - 1]减一，strNum[i]赋值9，这样这个整数就是89。就可以很自然想到对应的贪心解法了。

想到了贪心，还要考虑遍历顺序，只有从后向前遍历才能重复利用上次比较的结果
*/