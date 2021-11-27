type LFUCache struct {
    size int
    cap  int
    cache map[int]*Node
    freqHash map[int]*DLList
    minfreq int
}

type DLList struct {
    head *Node
    tail *Node
}

type Node struct {
    key int
    val int
    freq int
    pre *Node
    next *Node
}

func initDLList() *DLList{
    head, tail := &Node{}, &Node{}
    head.next = tail
    tail.pre = head
    return &DLList{
        head:head,
        tail:tail,
    }
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        cap: capacity,
        cache: make(map[int]*Node),
        freqHash: make(map[int]*DLList),
    }
}


func (this *LFUCache) Get(key int) int {
    v, ok:= this.cache[key]
    if !ok {
        return -1
    }
    this.IncFreq(v)
    return v.val
}


func (this *LFUCache) Put(key int, value int)  {
    if this.cap == 0 {
        return 
    }
    v, ok:= this.cache[key]
    if ok {
        this.cache[key].val = value
        this.IncFreq(v)
    } else {
        //先判断满没满
        if this.size >= this.cap {
            last := this.freqHash[this.minfreq].RemoveLast()
            delete(this.cache, last.key)
            this.size--
        }
        newNode := &Node{ key:key, val:value, freq:1 }
        this.cache[key] = newNode
        if this.freqHash[1] == nil {
            this.freqHash[1] = initDLList()
        }
        this.freqHash[1].AddToHead(newNode)
        this.minfreq = 1
        this.size++
    }
}

func (this *LFUCache) IncFreq(node *Node) {
    tempFreq := node.freq
    // 现在原来对应的频次删掉 以及 minfreq是否增加
    this.freqHash[tempFreq].RemoveNode(node)
    if this.minfreq == tempFreq && this.freqHash[tempFreq].IsEmptyDLL() {
        this.minfreq++
        delete(this.freqHash, tempFreq)
    }
    node.freq++
    if this.freqHash[node.freq] == nil {
        this.freqHash[node.freq] = initDLList()
    }
    this.freqHash[node.freq].AddToHead(node)
}

func (this *DLList) AddToHead(node *Node) {
    node.next = this.head.next
    node.pre  = this.head
    this.head.next.pre = node
    this.head.next = node
}

func (this *DLList) RemoveNode(node *Node) {
    node.pre.next = node.next
    node.next.pre = node.pre
    
    //
    node.pre = nil
    node.next = nil
}

func (this *DLList) RemoveLast () *Node {
    if this.IsEmptyDLL() {
        return nil
    }
    node:= this.tail.pre
    this.RemoveNode(node)
    return node
}

func (this *DLList) IsEmptyDLL () bool {
    return this.head.next == this.tail
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */