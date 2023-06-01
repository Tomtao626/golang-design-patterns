package main

import "fmt"

// 抽象层

// Car 抽象的汽车层
type Car interface {
	Run()
}

type Driver interface {
	Drive(c Car)
}

// 实现层

// BenZ 奔驰类
type BenZ struct{}

func (b *BenZ) Run() {
	fmt.Println("BenZ is Running")
}

// BmW 宝马类
type BmW struct{}

func (b *BmW) Run() {
	fmt.Println("BmW is Running")
}

// ZhangSan 张三车
type ZhangSan struct{}

func (z *ZhangSan) Drive(c Car) {
	fmt.Println("ZhangSan Drive Car")
	c.Run()
}

// LiSi 李四开车
type LiSi struct{}

func (l *LiSi) Drive(c Car) {
	fmt.Println("LiSi Drive Car")
	c.Run()
}

// 新增

// WangWu 王五开车
type WangWu struct{}

func (w *WangWu) Drive(c Car) {
	fmt.Println("WangWu Drive Car")
	c.Run()
}

func main() {
	// 张三既能开奔驰又能开宝马
	var bmw Car
	bmw = &BmW{}

	var benz Car
	benz = &BenZ{}

	var zhangSan Driver
	zhangSan = &ZhangSan{}
	zhangSan.Drive(bmw)
	zhangSan.Drive(benz)

	// 李四既能开奔驰又能开宝马
	var liSi Driver
	liSi = &LiSi{}
	liSi.Drive(benz)
	liSi.Drive(bmw)

	// 新增 王五既能开奔驰又能开宝马
	var wangWu Driver
	wangWu = &WangWu{}
	wangWu.Drive(bmw)
	wangWu.Drive(benz)
}
