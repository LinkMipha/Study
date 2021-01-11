package main


type LRUCache struct {
	size, capacity int
	cache map[int]*Lru
	head, tail *Lru
}

type Lru struct{
	key,value int
	pre,next *Lru
}

func initLru(ke,val int)*Lru{
	return &Lru{
		key:ke,
		value:val,
	}
}

func Constructor(capacity int) LRUCache {
	l:=LRUCache{
		cache : map[int]*Lru{},
		capacity:capacity,
		head:initLru(0,0),
		tail:initLru(0,0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	if _,ok:=this.cache[key];!ok{
		return -1
	}
	node:=this.cache[key]
	this.MoveHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok:=this.cache[key];!ok{
		this.size++
		node:=initLru(key,value)
		this.cache[key] = node
		this.AddToHead(node)
		if this.size>this.capacity{
			remove:=this.DeleteTail()
			delete(this.cache,remove.key)
			this.size--
		}
	}else{
		node := this.cache[key]
		node.value = value
		this.MoveHead(node)
	}
}

func (this *LRUCache)AddToHead(st *Lru){
	st.next = this.head.next
	this.head.next.pre = st
	this.head.next = st
	st.pre = this.head
}

func (this *LRUCache) reMove(node *Lru){
	node.pre.next = node.next
	node.next.pre = node.pre
}


func (this *LRUCache)MoveHead(st *Lru){
	this.reMove(st)
	this.AddToHead(st)
}
func (this *LRUCache)DeleteTail()*Lru{
	node:=this.tail.pre
	this.reMove(node)
	return node
}

func main()  {

}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */