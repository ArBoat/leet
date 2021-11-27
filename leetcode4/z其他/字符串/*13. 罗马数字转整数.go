var m = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
    n:=len(s)
    res:=0
    for i:= range s {
		注意边界
        if i < n-1 && m[s[i]] < m[s[i+1]] {
            res-=m[s[i]]
        } else {
            res+=m[s[i]]
        }
    }
    return res
}

// 左边比右边小减去， 大，加上