type LRUCache struct {
    size int
    capacity int
    cache map[int]*DLList
    head *DLList
    tail *DLList
}

type DLList struct {
    key int
    value int 
    pre *DLList
    next *DLList
}

func Constructor(capacity int) LRUCache {
    l:=LRUCache{
        capacity: capacity,
        cache: make(map[int]*DLList),
        head: &DLList{key:0, value:0},
        tail: &DLList{key:0, value:0},
    }
    // IM: l.
    l.head.next = l.tail
    l.tail.pre  = l.head
    return l
}


func (this *LRUCache) Get(key int) int {
    v, ok:= this.cache[key]
    if !ok {
        return -1
    }
    this.moveToHead(v)
    return v.value
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok:= this.cache[key]
    if !ok {
        // IM: new
        newNode:=&DLList{key:key, value:value}
        // 同步cache, =newNode
        this.cache[key]=newNode
        this.addToHead(newNode)
        this.size++
        if this.size > this.capacity {
            last:=this.removeLast()
            delete(this.cache,last.key)
            this.size--
        }
    } else {
        node.value = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) addToHead(node *DLList) {
    node.pre = this.head
    // IM: this.head.next
    node.next = this.head.next
    this.head.next.pre = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *DLList) {
    node.next.pre = node.pre
    node.pre.next = node.next
}

func (this *LRUCache) moveToHead(node *DLList) {
    // 一定是先删再加
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeLast() *DLList {
    node:=this.tail.pre
    this.removeNode(node)
    return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */