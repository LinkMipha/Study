package main


import (
	"fmt"
)

type Sayer interface {
	Say(message string)
	SayHi()
}

type Animal struct {
	Name string
}

func (ani *Animal)Say(message string)  {
	fmt.Printf("animal %v say %v",ani.Name,message)
}

func (a *Animal) SayHi() {
	a.Say("Hi")
}

type Dog struct {
	Animal
}

func (ani * Dog)Say(message string)  {
	fmt.Printf("animal %v say %v",ani.Name,message)
}


func inteface() {

	var sayer Sayer
	sayer = &Dog{Animal{Name: "Yoda"}}//子类未实现sayhi调用父类
	sayer.Say("hello world")
	sayer.SayHi()

}


type inter interface {
	funone()
}

type Userone struct {

}

func (Userone)funone()  {
	fmt.Println("test")
}

func main()  {
	var one *Userone
	if one==nil{
		fmt.Println("one is nil")
	}else {
		fmt.Println("one isnot nil")
	}

	var two inter
	two = one
	if two==nil{
		fmt.Println("two is nil")
	}else {
		fmt.Println("two isnot nil")
	}
}


