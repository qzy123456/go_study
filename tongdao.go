package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, (i+1)*100)
    }
}
func say2(s string, ch chan int) {
    for i := 0; i < 5; i++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Println(s, (i+1)*150)
    }
    ch <- 0
    close(ch)
}

func main() {
    ch := make(chan int)
    var c1  <-chan  int = ch
    go say2("world", ch)
    say("hello")
    fmt.Println(<-ch)
    fmt.Println("c1 is", <-c1)
}
