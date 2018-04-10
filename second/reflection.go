package main

import (
	"fmt"
	"reflect"
)

type IUser interface {
	//定义用户说话接口
	say(something string)
}

type User struct {
	//用户信息结构
	name string
	age  int
	sex  int
}

func (user User) say(something string) {
	fmt.Println(something);
}

func main() {
	user := User{"zhangsan", 10, 0}
	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)
	fmt.Println("Tname:", t.Name())
	fmt.Println("Vname:",v.Kind())

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i)
		fmt.Println()
		fmt.Printf("%s:%v = %v", f.Name, f.Type, val)
	}
}
