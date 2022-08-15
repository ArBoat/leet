func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

/*
F(n) = F(n - 1) + F(n - 2)
0,1,1,2,3,5
*/