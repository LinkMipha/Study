package main

import "sync"

//安全关闭channel

type Mychannel struct {
	C chan int
	Close bool
	Mute sync.Mutex
}

func GetChannel()*Mychannel  {
	return &Mychannel{C: make(chan int)}
}

func (this*Mychannel)SafeClose()  {
	this.Mute.Lock()
	if !this.Close{
		close(this.C)
		this.Close = true
	}
	this.Mute.Unlock()
}

//判断是否关闭
func (this *Mychannel)Isclose()bool  {
	this.Mute.Lock()
	defer this.Mute.Unlock()
	return this.Close
}


type OnceChannel struct {
	Ch chan int
	once sync.Once
}

//使用once关闭

func GetOnceChannel()*OnceChannel  {
	return &OnceChannel{Ch: make(chan int)}
}


func (this*OnceChannel)SafeCloseByOnce()  {
	this.once.Do(func() {
		close(this.Ch)
	})
}


func main()  {

}
