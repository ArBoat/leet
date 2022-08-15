func canCompleteCircuit(gas []int, cost []int) int {
	curSum, totalSum := 0, 0
	start := 0
	for i := range gas {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

/*
1: totalSum =  toalGas-totalCost  > 0
2: 由于解存在而且唯一，curSum >0 第一个位置就是答案

O(n)/O(1)
*/