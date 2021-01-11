package main

//存放的数据为Link
type Link struct {
	//内部数据
}

//map key放置数据
type Set struct {
	m map[Link]bool
}

//获取Set
func GetSet() *Set {
	return &Set{
		m: map[Link]bool{},
	}
}

//返回key的数组
func (s *Set)GetArray()*[]Link  {
	len :=len(s.m)
	array:=make([]Link,len)
	sum:=0
	for k,_:=range s.m{
		array[sum] = k
		sum++
	}
	return &array
}


//增加元素，判断是否存在
func (s*Set)Add(link Link)bool  {
	if _,ok:=s.m[link];ok{
		return false
	}else {
		s.m[link] = true
		return true
	}
}

//删除元素
func (s *Set)Remove(link Link)  {
	 delete(s.m,link)
}

//清除元素
func (s *Set)Clear()  {
	s.m  = make(map[Link]bool)
}

//获取长度
func (s *Set)GetLen()int  {
	return len(s.m)
}



