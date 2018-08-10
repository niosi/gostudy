package main

import (
    "fmt"
    "time"
)

func main2() {
    for range time.Tick(3 * time.Second) {
        fmt.Println(time.Now().Format(time.ANSIC))
    }
}