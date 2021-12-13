// 快慢指针
// 判断成环问题
func isHappy(n int) bool {
	slow, fast := n, getSum(n)
	for fast != 1 && slow != fast {
		slow = getSum(slow)
		fast = getSum(getSum(fast))
	}
	return fast == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10) // %取余数， /除法
		n /= 10
	}
	return sum
}

// hash
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true //先存hash，再取sum
		n = getSum(n)
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10) // %取余数， /除法
		n /= 10
	}
	return sum
}