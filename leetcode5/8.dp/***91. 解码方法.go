func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1 //IM : = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		//IM : i>=2
		if i >= 2 && s[i-2] != '0' && 10*(s[i-2]-'0')+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

/*
爬楼梯变种
dp[i] = dp[i-1]+dp[i-2]
dp[0] = 1 //空字符串可以有 1 种解码方法，解码出一个空字符串


a:= s[i-1]-'0'
b:= 10*(s[i-2]-'0') + (s[i-1]-'0')
dp[i] = dp[i-1] dp[i]!='0'   1<=a<=9
dp[i] += dp[i-2]             10<=b<=26
*/