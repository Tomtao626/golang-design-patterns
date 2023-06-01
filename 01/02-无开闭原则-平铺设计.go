package main

import "fmt"

// Banker 抽象的银行业务员
type Banker struct{}

// Save 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了 存款业务...")
}

// Transfer 转账业务
func (b *Banker) Transfer() {
	fmt.Println("进行了 转账业务...")
}

// Pay 支付业务
func (b *Banker) Pay() {
	fmt.Println("进行了 支付业务...")
}

// Stock 股票业务
func (b *Banker) Stock() {
	fmt.Println("进行了 股票业务...")
}

func main() {
	b := &Banker{}
	b.Save()
	b.Transfer()
	b.Pay()
	// 新业务加进来
	b.Stock()
}

// 如果银行业务员有新的业务Stock(股票)进来, 需要对Banker类进行修改 破坏了原有类的结构 增强了耦合性 是不倡导的
