func isLongPressedName(name string, typed string) bool {
	i := 0
	for j := 0; j < len(typed); j++ {
		if i < len(name) && name[i] == typed[j] {
			i++
		} else if j > 0 && typed[j] == typed[j-1] {
			continue
		} else {
			return false
		}
	}
	return i == len(name)
}

/*
双指针：
typed和name都要遍历到最后一个字符，但typed长度要大于等于name长度
外层typed每次+1，内层name有条件+1(保证了typed>=name的长度)
j>0 既保证了数组不会越界，也保证了第一个字母不相等时候直接返回false
*/