package design_pattern

import "sync"

//需要的类
type Link struct {
	value int
}

var link *Link

//未加锁懒汉模式
func GetInstance()*Link  {
	if link==nil{
		link = new(Link)
	}
	return link
}


//加锁，为了线程安全
var lock sync.Mutex
//线程安全的懒汉模式
func GetInstanceSync()*Link{
	defer lock.Unlock()
	lock.Lock()
	if link==nil{
		link = new(Link)
	}
	return link
}


var Mipha Link
//饿汉式,直接建立实体，浪费内存
func GetInstanceStandard() *Link{
	return &Mipha
}


//双重检查
func GetTwoSelect()*Link {
	if link==nil{
		lock.Lock()

		//再判断
		if link==nil{
			link  = new(Link)
		}

		lock.Unlock()
	}
	return link
}

var once sync.Once
//once方式建立
func GetOnce() *Link {
	once.Do(func() {
		link  = new(Link)
	})
	return link
}

func main()  {

}
