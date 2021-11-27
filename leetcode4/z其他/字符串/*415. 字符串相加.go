func addStrings(num1 string, num2 string) string {
    res:=""
    carry:=0
    for i, j := len(num1)-1, len(num2)-1; i>=0 || j>=0 || carry != 0; i, j = i-1, j-1 {
        x, y:= 0,0
        if i >= 0 {
            x = int(num1[i]-'0')
        }
        if j >= 0 {
            x = int(num2[j]-'0')
        }
        temp := x + y + carry
        res = strconv.Itoa(temp%10) + res
        carry = temp/10
    }
    return res
}