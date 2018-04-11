package main

import (
	"fmt"
	"runtime"
)

func main() { //goroutine
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int,10)
	for i := 0; i <= 10; i++ {
		go sum(ch, i)
	}
	for i := 0; i <= 10; i++ {
		<-ch
	}
	/*go Go(ch)
	for v := range ch{
		fmt.Print(v)
	}*/
}

func Go(a chan int) {
	fmt.Println("run start")
	a <- 0
	close(a)
}

func sum(a chan int, index int) {
	sum := 0
	for i := 0; i <= 1000000; i++ {
		sum += i
	}
	fmt.Println(index, sum)
	a <- 0
}
