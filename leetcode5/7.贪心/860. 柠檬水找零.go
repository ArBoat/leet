func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
		} else if v == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 2 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

/*
情况一：账单是5，直接收下。
情况二：账单是10，消耗一个5，增加一个10
情况三：账单是20，优先消耗一个10和一个5，如果不够，再消耗三个5
*/