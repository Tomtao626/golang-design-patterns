package main

import "fmt"

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// CatB 给小猫添加一个可以睡觉的方法
type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// CatC 给小猫添加一个可以睡觉的方法(采用组合的方式)
type CatC struct {
	C *Cat // 组合进来一个Cat类
}

func (cc *CatC) Eat() {
	cc.C.Eat() // 调用组合类的功能
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

type CatD struct{}

func (cd *CatD) Eat(c *Cat) {
	c.Eat() // 接口依赖
}

func (cd *CatD) Sleep(c *CatC) {
	c.Sleep() // 接口依赖
}

func main() {
	c := &Cat{}
	c.Eat()

	fmt.Println("--------")
	cb := &CatB{}
	cb.Eat()
	cb.Sleep()

	fmt.Println("--------")
	cc := &CatC{}
	cc.Eat()
	cc.C.Eat()
	cc.Sleep()

	fmt.Println("--------")
	cd := &CatD{}
	cd.Eat(c)
	cd.Sleep(cc)
}

// 出现继承和复合都能实现的情况 优先使用组合
