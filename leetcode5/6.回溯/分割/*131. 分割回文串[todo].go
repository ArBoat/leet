func partition(s string) [][]string {
	res := [][]string{}
	path := []string{}
	var backtrack func(path []string, cur int)
	backtrack = func(path []string, cur int) {
		if cur == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := cur; i < len(s); i++ {
			if !isWord(s, cur, i) {
				continue
			}
			backtrack(append(path, s[cur:i+1]), i+1) //IM:i+1
		}
	}
	backtrack(path, 0)
	return res

}

func isWord(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*
O(n*2^n)
O(m^2)
*/