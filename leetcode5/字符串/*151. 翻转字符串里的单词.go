func reverseWords(s string) string {
	b := []byte(s)
	if len(b) == 0 {
		return s
	}
	slow, fast := 0, 0
	//1.1 delete head blank
	for fast < len(b) && b[fast] == ' ' {
		fast++
	}
	//1.2 delete body blank
	for ; fast < len(b); fast++ {
		if fast-1 > 0 && b[fast] == b[fast-1] && b[fast] == ' ' {
			continue
		}
		b[slow] = b[fast]
		slow++ // slow指向下一个不是重复的位置（空格会被保留一个）
	}
	//1.3 delete tail blank 因为最后一个空格会被保留
	if slow-1 > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}
	// 2.reverse whole string
	reverse(b, 0, len(b)-1)
	// 3. reverse single world
	left := 0
	for left < len(b) {
		right := left
		for right < len(b) && b[right] != ' ' {
			right++
		}
		reverse(b, left, right-1)
		left = right + 1
	}
	return string(b)
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++ //IM: do not forget
		right--
	}
}

/*
1.双指针删除冗余的前中后
2.翻转整个字符串
3.翻转单词
*/