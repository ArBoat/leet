type MyHashMap struct {
    hash [769]*Link
}

type Link struct {
    key int
    value int
    next *Link
}

func getHashMod(key int) int {
	//映射
    return key%769
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    v := getHashMod(key)
	// 不存在映射
    if this.hash[v] == nil {
        this.hash[v] = &Link{
            key:key,
            value:value,
        }
        return
    }
    p:=this.hash[v]
    for {
		//存在
        if p.key == key {
            p.value = value
            return
        }
		//不存在
        if p.next == nil {
            p.next = &Link{
                key:key,
                value:value,
            }
            return
        }
        p=p.next
    }
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    v := getHashMod(key)
    p:=this.hash[v]
	//找
    for p!=nil {
        if p.key == key {
            return p.value
        }
        p=p.next
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    v := getHashMod(key)
	// 不存在
    if this.hash[v] == nil {
        return
    }
	// 第一个
    if this.hash[v].key == key {
        this.hash[v] = this.hash[v].next
        return
    }
	// 不是第一个
    pre:=this.hash[v]
    p := pre.next
    for p!=nil {
        if p.key == key {
            pre.next = p.next
            return
        }
        pre = p
        p = p.next
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */