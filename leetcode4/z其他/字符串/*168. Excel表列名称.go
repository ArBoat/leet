// 字符串问题想好 byte string int转换，

// byte
func convertToTitle(columnNumber int) string {
    res:=[]byte{}
    for columnNumber > 0 {
        columnNumber--
        res = append([]byte{byte(columnNumber%26)+'A'},res...)
        columnNumber/=26
    }
    return string(res)
}

// byte 后排序
func convertToTitle(columnNumber int) string {
    res:=[]byte{}
    for columnNumber > 0 {
        columnNumber--
        res = append(res,byte(columnNumber%26)+'A')
        columnNumber/=26
    }
    for i, n := 0, len(res); i < n/2; i++ {
        res[i], res[n-1-i] = res[n-1-i], res[i]
    }
    return string(res)
}

// string
func convertToTitle(columnNumber int) string {
    res:=[]string{}
    for columnNumber > 0 {
        columnNumber--
        res = append([]string{string(byte(columnNumber%26)+'A')},res...)
        columnNumber/=26
    }
    return strings.Join(res,"")
}