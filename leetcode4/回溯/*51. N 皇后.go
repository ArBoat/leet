func solveNQueens(n int) [][]string {
    res:= [][]string{}
    path:= make([]int, n)
    for i:=range path {
        path[i]=-1
    }
    c, d1, d2 := make(map[int]bool),make(map[int]bool),make(map[int]bool)
    var backtrack func(path []int, row int, c, d1, d2 map[int]bool)
    backtrack = func(path []int, row int, c, d1, d2 map[int]bool){
        // end
        
        if row == n {
            board := genetate(path, n)
            res = append(res, board)
            return
        }
        for i:=0; i<n ; i ++ {
            if c[i] || d1[row-i] || d2[row+i]{
                continue
            }
            path[row]=i
            c[i],d1[row-i],d2[row+i] = true, true, true
            backtrack(path, row+1, c, d1, d2)
            path[row] = -1
            delete(c, i)
            delete(d1, row-i)
            delete(d2, row+i)
        }
    }
    backtrack(path, 0, c, d1, d2)
    return res
}

func genetate(path []int, n int) []string {
    res:=[]string{}
    for i:=0;i<n;i++{
        row:=make([]byte, n)
        for j:= range row {
            row[j] = '.'
        }
        row[path[i]] = 'Q'
        res = append(res, string(row))
    }
    return res
}
/*
* 时间复杂度O(N!)
* 空间复杂度O(N)
backtrack
*/