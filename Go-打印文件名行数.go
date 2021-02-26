package main

import (
    "runtime"
    "fmt"
)

func trace(s string) {
        pc := make([]uintptr, 10)  
        runtime.Callers(2, pc)
        f := runtime.FuncForPC(pc[0])
        file, line := f.FileLine(pc[0])
        fmt.Printf("%s:%d %s -- %s\n", file, line, f.Name(), s) //这里可以换成自带的 log
}

func foo(s  string){
        trace(s)
}

func main(){
    foo("test")
}
