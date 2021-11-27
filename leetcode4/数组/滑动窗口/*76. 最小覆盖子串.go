func minWindow(s string, t string) string {
    win, hash:= make(map[rune]int), make(map[rune]int)
    left, right := 0, 0
    match, size:= 0, len(s)+1
    rs:=[]rune(s)
    i:=0
    for _, v:= range t{
        hash[v]++
    }
    for j, v := range rs {
        if _, ok:=hash[v];ok{
            win[v]++
            if win[v] == hash[v] {
                match++
            }
        }
        for ;match == len(hash);i++ {
            if j-i+1 < size {
                size = j-i+1
                left, right = i, j
            }
            temp:=rs[i]
            if _, ok:=hash[temp]; ok{
                win[temp]--
                if win[temp] < hash[temp]{
                    match--
                }
            }
        }
    }
    if size == len(s)+1 {
        return ""
    } 
    return string(rs[left:right+1])
}