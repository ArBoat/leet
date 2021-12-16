func hammingDistance(x int, y int) int {
    v:=x^y
    res:=0
    for v != 0 {
        v = v & (v-1)
        res++
    }
    return res
}