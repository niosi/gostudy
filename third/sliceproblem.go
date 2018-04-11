package main

import "fmt"

func appendNum(s []int)  []int {
	s = append(s, 3)
	return s
}
func main() {
	s := make([]int,0)
	fmt.Println(s)
	fmt.Println(appendNum(s))
}
