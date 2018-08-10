package main
import (
	"fmt"
   "time"

)

func enqueue(q chan int) {
   time.Sleep(time.Second * 3)
    q <- 10
    close(q)
}

func main() {
    var c =make(chan int)
    go enqueue(c)
    for {
        select {
            case x, ok := <- c:
               if ok {
                   fmt.Println(x)
               } else {
                   fmt.Println("closed")
                   return
               }
           default:
               fmt.Println("waiting")
               time.Sleep(time.Second)
        }
    }
}