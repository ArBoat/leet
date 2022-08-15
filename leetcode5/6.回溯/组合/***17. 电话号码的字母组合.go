func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}
	hash := make(map[string][]string)
	hash["2"] = []string{"a", "b", "c"}
	hash["3"] = []string{"d", "e", "f"}
	hash["4"] = []string{"g", "h", "i"}
	hash["5"] = []string{"j", "k", "l"}
	hash["6"] = []string{"m", "n", "o"}
	hash["7"] = []string{"p", "q", "r", "s"}
	hash["8"] = []string{"t", "u", "v"}
	hash["9"] = []string{"w", "x", "y", "z"}

	var backtrack func(path string, cur int)
	backtrack = func(path string, cur int) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}
		// string index byte
		letters := hash[string(digits[cur])]
		for _, v := range letters {
			backtrack(path+v, cur+1)
		}
	}
	backtrack("", 0)
	return res
}

/*
O(3^m*4^n),3个4个字母对应的按键
O（m+n）
*/