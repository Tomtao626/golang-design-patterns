package main

import "fmt"

// 职责划分清楚 对外只提供一种功能,而引起类变化的原因只有一个

// ClothesShop 逛街穿衣服
type ClothesShop struct{}

func (c *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

// ClothesWork 工作穿衣服
type ClothesWork struct{}

func (c *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {
	// 工作的业务
	cw := &ClothesWork{}
	cw.Style()

	// 逛街的业务
	cs := &ClothesShop{}
	cs.Style()
}

/*
// Clothes 穿衣服的方式
type Clothes struct{}

// OnWork 工作穿衣服
func (c *Clothes) OnWork() {
	println("工作的装扮")
}

// OnShop 逛街穿衣服
func (c *Clothes) OnShop() {
	println("逛街的装扮")
}

func main() {
	c := &Clothes{}

	// 工作的业务
	fmt.Println("工作的业务")
	c.OnWork()

	// 逛街的业务
	fmt.Println("逛街的业务")
	c.OnShop()
}
*/
