package main

import "fmt"

// Benz 奔驰汽车
type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("Benz is running")
}

// Bmw 宝马汽车
type Bmw struct{}

func (b *Bmw) Run() {
	fmt.Println("Bmw is running")
}

// Zhang3 司机张三
type Zhang3 struct{}

func (z *Zhang3) DriveBenz(benz *Benz) {
	fmt.Println("Zhang3 Drive Benz")
	benz.Run()
}

func (z *Zhang3) DriveBmw(bmw *Bmw) {
	fmt.Println("Zhang3 Drive Bmw")
	bmw.Run()
}

// Li4 司机李四
type Li4 struct{}

func (l *Li4) DriverBmw(bmw *Bmw) {
	fmt.Println("Li4 Drive Bmw")
	bmw.Run()
}

func (l *Li4) DriverBenz(benz *Benz) {
	fmt.Println("Li4 Drive Benz")
	benz.Run()
}

func main() {
	// 业务1 张三开奔驰
	benz := &Benz{}
	zhang3 := &Zhang3{}
	zhang3.DriveBenz(benz)

	// 业务2 李四开宝马
	bmw := &Bmw{}
	li4 := &Li4{}
	li4.DriverBmw(bmw)
}

// 不论是宝马还是奔驰 都有Run方法 而李四和张三都有驾驶方法 应该对应的抽象出来 作为公共的抽象层
