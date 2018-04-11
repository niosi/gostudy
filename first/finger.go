package main

import "fmt"
var a string = "ss"
var f1 *string
var f2 **string


func main() {
	f1 = &a;
	f2 = &f1;
	fmt.Printf("原始数据:%s\n", a)
	fmt.Printf("指针1的数据:%s\n", *f1)
	fmt.Printf("指针2的数据:%s\n", **f2)
}
