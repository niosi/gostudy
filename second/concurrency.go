package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() { //goroutine
	runtime.GOMAXPROCS(runtime.NumCPU())
	//ch := make(chan int,10)
	//for i := 0; i <= 10; i++ {
	//	go sumbychannel(ch, i)
	//}
	//for i := 0; i <= 10; i++ {
	//	<-ch
	//}
	/*go Go(ch)
	for v := range ch{
		fmt.Print(v)
	}*/
	wg := sync.WaitGroup{}
	wg.Add(11)
	for i := 0; i <= 10; i++ {
		go sum2bysync(&wg, i)
	}
	wg.Wait()
}

func Go(a chan int) {
	fmt.Println("run start")
	a <- 0
	close(a)
}

func sumbychannel(a chan int, index int) {
	sum := 0
	for i := 0; i <= 1000000; i++ {
		sum += i
	}
	fmt.Println(index, sum)
	a <- 0
}

func sum2bysync(wg *sync.WaitGroup,index int)  {
	sum := 0
	for i := 0; i <= 1000000; i++ {
		sum += i
	}
	fmt.Println(index, sum)
	wg.Done()
}
