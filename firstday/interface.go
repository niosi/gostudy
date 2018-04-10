package main

import (
	"fmt"
)

type Phone interface {//type 为 interface
	//定义接口
	call()
	selfFun()
}

type NokiaPhone struct {
	a int
	b string
}

func (nokiaPhone NokiaPhone) selfFun() {
	fmt.Println("Nokia My Self Fun" + nokiaPhone.b);
}

func (nokiaPhone NokiaPhone) call() { //实现接口
	fmt.Println("I am Nokia, I can call you!")
}

func (testPhone *IPhone)  test(){
	fmt.Println("TEST FUN");
}

type IPhone struct {//type 为 struct
	a int
	b string
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func (iPhone IPhone) selfFun() {
	fmt.Println("iPhone My Self Fun" + iPhone.b);
}

type MiPhone func(int, int) //定义类型type 为fun

func (miPhone MiPhone) myfun() { //为类型定义私有方法
	fmt.Println("mi myfun");
}

func main() {
	nphone := NokiaPhone{a: 3, b: " Hello Nokia"}
	nphone.call()
	nphone.selfFun()

	iphone := new(IPhone)
	iphone.call()
	iphone.b = " Hello IPhone"
	iphone.selfFun()
	(*IPhone).test(&IPhone{})//这种写法含义
	iphone2 := IPhone{}
	iphone2.test()

	miPhone := MiPhone(func(a int, b int) {
		fmt.Printf("a + b = %d", a+b)
	})
	miPhone.myfun()
	miPhone(1, 1)
}
