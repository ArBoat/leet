func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel
	//i for  station
	for i := 0; i < n; i++ { // im: i:=0 ,<n
		for t := i; t >= 0; t-- {
			if dp[t] >= stations[i][0] {
				dp[t+1] = max(dp[t+1], dp[t]+stations[i][1])
			}
		}
	}
	for i, v := range dp {
		if v >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
dp[i]为i次油最远距离， 求dp[i]>=target的最小i
dp[0] =start
dp[t+1] = max(dp[t+1], dp[t]+stations[i][1])
*/