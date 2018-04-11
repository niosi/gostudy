package main

import "fmt"
import (
	"reflect"
)

type User struct {
	uname string
	uage int
	usex int
}

type Student struct {
	User
	sno string
}

func (user User) sayHello(str string)  {
	fmt.Printf("My Name Is :%s\n I Say %s\n",user.uname,str)
}

func (stu Student) sayHello(str string)  {
	fmt.Printf("My Name Is :%s\n My Sno is %d\n I Say %s\n",stu.User.uname,stu.sno,str)
}

func main() {
	user := User{uname:"张三",uage:18,usex:1}
	stu := Student{User:User{"李四",18,0},sno:"11001100"}
	IInfo(user)
	IInfo(stu)
	IInfo(nil)
}

func IInfo(i interface{})  {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	switch i.(type) {
	case User:
		fmt.Println("isUser");
		fmt.Println(t.FieldByIndex([]int{0}))
		fmt.Println(v.FieldByIndex([]int{0}))
	case Student:
		fmt.Println("isStudent");
		fmt.Println(t.FieldByIndex([]int{0,0}))
		fmt.Println(v.FieldByIndex([]int{0,0}))
	default:
		fmt.Println("unknowtype")
	}
}
