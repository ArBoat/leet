func restoreIpAddresses(s string) []string {
	res := []string{}
	path := []string{}
	var backtrack func(path []string, index int)
	backtrack = func(path []string, index int) {
		// IM: index == len(s)
		if index == len(s) && len(path) == 4 {
			ip := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
			res = append(res, ip)
			return
		}
		for i := index; i < len(s); i++ {
			//字符小于3， 长度小于4， valid
			if i-index < 3 && len(path) < 4 && isValid(s, index, i) {
				backtrack(append(path, s[index:i+1]), i+1)
			}
		}
	}
	backtrack(path, 0)
	return res
}

//[]
func isValid(s string, start, end int) bool {
	// pre "0"
	if end-start > 0 && s[start] == '0' {
		return false
	}
	num, err := strconv.Atoi(s[start : end+1])
	if err != nil {
		return false
	}
	if num > 255 {
		return false
	}
	return true
}