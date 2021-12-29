func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	hash := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && hash[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

/*
hash： 右：左
遍历
左括号直接加stack
len(stack)>0且右括号等于栈顶：出栈
否则false

len（stack） == 0
*/