//原地倒序扩容
func replaceSpace(s string) string {
	b := []byte(s)
	originLen := len(b)
	if originLen == 0 {
		return s
	}
	blankCount := 0
	for i := range b {
		if b[i] == ' ' {
			blankCount++
		}
	}
	if blankCount == 0 {
		return s
	}
	newLen := originLen + blankCount*2
	res := make([]byte, newLen)
	cur := newLen - 1
	for i := originLen - 1; i >= 0; i-- {
		if b[i] == ' ' {
			res[cur] = '0'
			cur--
			res[cur] = '2'
			cur--
			res[cur] = '%'
			cur--
		} else {
			res[cur] = b[i]
			cur--
		}
	}
	return string(res)
}

// 遍历添加
func replaceSpace(s string) string {
	b := []byte(s)
	res := make([]byte, 0)
	for i := range b {
		if b[i] == ' ' {
			res = append(res, []byte("%20")...)
		} else {
			res = append(res, b[i])
		}
	}
	return string(res)
}