package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Name string
    Age  int
    Addr string
}

func NewUser(name string, age int) User {
    u := User{}
    u.Name = name
    u.Age = age
    return u
}

func main1() {
    u1 := NewUser("user1", 12)
    u2 := NewUser("user1", 12)
    fmt.Println(reflect.DeepEqual(u1, u2))
    fmt.Println(reflect.DeepEqual([]int{1, 2}, []int{1, 2}))
    fmt.Println(reflect.DeepEqual([2]int{1, 2}, [2]int{1, 2}))
    fmt.Println(reflect.DeepEqual(map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 2}))
}