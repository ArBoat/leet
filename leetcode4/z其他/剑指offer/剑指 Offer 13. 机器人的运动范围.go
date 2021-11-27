func movingCount(m int, n int, k int) int {
    dir:=[][]int{{0,1},{0,-1},{1,0},{-1,0}}

    visited:=make([][]bool, m)
    for i:= range visited {
        visited[i] = make([]bool, n)
    }

    count:=0

    var backtrack func(i,j int)
    backtrack = func(i,j int){
        if i>=m || i<0 || j>=n || j<0 || !valid(i,j,k) || visited[i][j]{
            return
        }
        count++
        visited[i][j] = true
        for _, v:= range dir {
            ni:=i+v[0]
            nj:=j+v[1]
            backtrack(ni, nj)
        }
        // 不需要恢复
    }

    backtrack(0,0)
    return count
}

func valid(m,n,k int) bool {
    return (sum(m) + sum(n))<=k
}

func sum(n int) int {
    res:=0
    for n != 0 {
        res += n%10
        n = n/10
    }
    return res
}