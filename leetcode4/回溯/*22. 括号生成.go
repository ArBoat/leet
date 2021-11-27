func generateParenthesis(n int) []string {
    if n == 0 {
        return nil
    }
    res:=make([]string, 0)
    var generate func(l, r int, s string)
    generate = func(l, r int, s string){
        if l == n && r == n {
            res = append(res, s)
        }
        if l<n{
            generate(l+1,r, s+"(")
        }
        if r<l{
            generate(l,r+1, s+")")
        }
    }
    generate(0,0,"")
    return res
}