package main

import "fmt"

type Fruit interface {
	Show()
}
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是梨子")
}

type Peach struct {
	// ...
}

func (peach *Peach) Show() {
	// ....
}

// Factory /* 工厂模块 */
// 一个工厂, 有一个生产水果的机器, 返回一个抽象水果的指针
type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	} else if kind == "peach" {
		fruit = new(Banana)
	}
	return fruit
}

// 业务逻辑层
func main() {
	factory := new(Factory)
	apple := factory.CreateFruit("apple")
	apple.Show() // 多态  苹果对象的方法

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
