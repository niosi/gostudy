package main

import (
    "testing"
    "fmt"
)

func TestDataStruct(t *testing.T) {
    for i := 0; i < 100000000; i++ {
        var a struct {
            Name string
            Age  int
        }
        a.Name = "Hello"
        a.Age = 10
    }
}

func TestDataMap(t *testing.T) {
    for i := 0; i < 100000000; i++ {
        var a = map[string]interface{}{}
        a["Name"] = "Hello"
        a["Age"] = 10
    }
}

func TestDiv(t *testing.T) {
    for i:=0;i<100000000;i++{
         fmt.Sprintf("%d", 100000000000/50)
    }
}

func TestDiv2(t *testing.T) {
    for i:=0;i<100000000;i++{
         fmt.Sprintf("%d", 100000000000*0.02)
    }
}