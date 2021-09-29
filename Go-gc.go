package main

import (

    "log"

    "runtime"

)
//gcTriggerHeap：当所分配的堆大小达到阈值（由控制器计算的触发堆的大小）时，将会触发。
//gcTriggerTime：当距离上一个 GC 周期的时间超过一定时间时，将会触发。-时间周期以 runtime.forcegcperiod 变量为准，默认 2 分钟。
//gcTriggerCycle：如果没有开启 GC，则启动 GC。在手动触发的 runtime.GC 方法中涉及。
var intMap map[int]int

var cnt = 8192

func main() {

    printMemStats()

    initMap()

    runtime.GC()

    printMemStats()

    log.Println(len(intMap))

    for i := 0; i < cnt; i++ {

        delete(intMap, i)

    }

    log.Println(len(intMap))

    runtime.GC()

    printMemStats()

    intMap = nil

    runtime.GC()

    printMemStats()

}

func initMap() {

    intMap = make(map[int]int, cnt)

    for i := 0; i < cnt; i++ {

        intMap[i] = i

    }

}

func printMemStats() {

    var m runtime.MemStats

    runtime.ReadMemStats(&m)

    log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)

}

