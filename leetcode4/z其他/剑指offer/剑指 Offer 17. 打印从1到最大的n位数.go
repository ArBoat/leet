// 大数全排列
func printNumbers(n int) []string {
	res:=[]string{}
	var dfs func(index, digit int, num []byte)
	dfs = func(index, digit int, num []byte){
		if index == digit{
			res = append(res, string(num))
			return
		}
		for i:=byte('0');i<=byte('9');i++{
			num[index] = i
			dfs(index+1,digit, num)
		}
	}
	for digit:=1;digit<n+1;digit++{
		for first :=byte('1'); first<=byte('9');first++{
			num :=make([]byte, digit)
			num[0] = first
			dfs(1, digit,num)
		}
	}
	return res
}


//简单
func printNumbers(n int) []int {
    num:=1
    for i:=1;i<=n;i++{
        num*=10
    }
    num = num-1
    res:= []int{}
    for i:=1;i<=num;i++{
        res= append(res, i)
    }
    return res
}