func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err != nil { // IM： 如果不是数字，则是操作符， 不要反
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		} else {
			stack = append(stack, val)
		}
	}
	return stack[0]
}