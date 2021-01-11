package main

import "fmt"

type Typec interface {
	UseTypec()string
}

type USB interface {
	UseUsb()string
}

func NewUsb()USB  {
	return &keyboard{}
}
type keyboard struct {
	
}

func (k *keyboard)UseUsb()string  {
	return fmt.Sprint("keyboard实现UseUsb")
}

func NewAdapter(keyboard USB)Typec  {
	return &Adapter{keyboard}
}

type Adapter struct {
	USB
}

func (a *Adapter)UseTypec() string {
	return fmt.Sprintf(a.UseUsb()+"Usb use Adapter to UseTypec")
}

func main()  {

	usb := NewUsb()
	adapter := NewAdapter(usb)
	fmt.Println(adapter.UseTypec())

}