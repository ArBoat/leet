//unicode
func isAnagram(s, t string) bool {
    if len(s) != len(t) {
        return false
    }
    cnt := map[rune]int{}
    for _, ch := range s {
        cnt[ch]++
    }
    for _, ch := range t {
        cnt[ch]--
        if cnt[ch] < 0 {
            return false
        }
    }
    return true
}

//
func isAnagram(s string, t string) bool {
    var n1, n2 [26]int
    for _, v:= range s {
        n1[v-'a']++
    }
    for _, v:= range t {
        n2[v-'a']++
    }
    return n1 == n2
}