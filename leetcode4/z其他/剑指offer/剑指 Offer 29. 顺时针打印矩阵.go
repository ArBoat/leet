// 同lc54


func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }
    m,n:=len(matrix), len(matrix[0])
    res:=make([]int, m*n)
    l, r, t, b:= 0, n-1, 0, m-1
    num:=0
    total:= m*n
    for num<total{
		// num<total
        // l to r
        for i:=l;i<=r && num<total;i++{
            res[num] = matrix[t][i]
            num++
        }
        t++
        // t to b
        for i:=t;i<=b && num<total;i++{
            res[num] = matrix[i][r]
            num++
        }
        r--
        // r to l
        for i:=r;i>=l && num<total;i--{
            res[num] = matrix[b][i]
            num++
        }
        b--
        // b to t
        for i:=b;i>=t && num<total;i--{
            res[num] = matrix[i][l]
            num++
        }
        l++
    }
    return res
}