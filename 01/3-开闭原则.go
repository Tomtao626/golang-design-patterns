package main

import "fmt"

// AbstractBanker 实现一个抽象的银行业务员
type AbstractBanker interface {
	DoSomething()
}

// SaveBanker 存款的业务员
type SaveBanker struct {
	//AbstractBanker
}

func (sb *SaveBanker) DoSomething() {
	fmt.Println("进行了 存款业务...")
}

// TransferBanker 转账的业务员
type TransferBanker struct {
	//AbstractBanker
}

func (tb *TransferBanker) DoSomething() {
	fmt.Println("进行了 转账业务...")
}

// PayBanker 支付的业务员
type PayBanker struct {
	//AbstractBanker
}

func (pb *PayBanker) DoSomething() {
	fmt.Println("进行了 支付业务...")
}

// StackBanker 股票的业务员
type StackBanker struct {
	//AbstractBanker
}

func (sb *StackBanker) DoSomething() {
	fmt.Println("进行了 股票业务...")
}

// BankerBusiness 实现一个架构层(基于抽象层进行封装-针对interface接口进行封装)
func BankerBusiness(banker AbstractBanker) {
	// 把接口作为函数的参数,可以动态的替换为接口的实现类,从而实现多态
	// 通过接口向下调用(多态的现象)
	banker.DoSomething()
}

func main() {
	/*
		// 存款的业务
		sb := SaveBanker{}
		sb.DoSomething()

		// 转账的业务
		tb := TransferBanker{}
		tb.DoSomething()

		// 支付的业务
		pb := PayBanker{}
		pb.DoSomething()

		// 股票的业务
		sb2 := StackBanker{}
		sb2.DoSomething()
	*/
	// 存款的业务
	BankerBusiness(&SaveBanker{})

	// 转账的业务
	BankerBusiness(&TransferBanker{})

	// 支付的业务
	BankerBusiness(&PayBanker{})

	// 股票的业务
	BankerBusiness(&StackBanker{})
}
