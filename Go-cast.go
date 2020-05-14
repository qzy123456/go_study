package main

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
)

func main() {
	// ToString
	fmt.Println(cast.ToString("leedarjun"))        // leedarjun
	fmt.Println(cast.ToString(8))                  // 8
	fmt.Println(cast.ToString(8.31))               // 8.31
	fmt.Println(cast.ToString([]byte("one time"))) // one time
	fmt.Println(cast.ToString(nil))                // ""

	var foo interface{} = "one more time"
	fmt.Println(cast.ToString(foo))                // one more time

	// ToInt
	fmt.Println(cast.ToInt(8))      // 8
	fmt.Println(cast.ToInt(8.31))   // 8
	fmt.Println(cast.ToInt("8"))    // 8
	fmt.Println(cast.ToInt(true))   // 1
	fmt.Println(cast.ToInt(false))  // 0

	var eight interface{} = 8
	fmt.Println(cast.ToInt(eight))  // 8
	fmt.Println(cast.ToInt(nil))    // 0
	now := time.Now()
	timestamp := 1579615973
	timeStr := "2020-01-21 22:13:48"

	fmt.Println(cast.ToTime(now))       // 2020-01-22 06:31:50.5068465 +0800 CST m=+0.000997701
	fmt.Println(cast.ToTime(timestamp)) // 2020-01-21 22:12:53 +0800 CST
	fmt.Println(cast.ToTime(timeStr))   // 2020-01-21 22:13:48 +0000 UTC

	d, _ := time.ParseDuration("1m30s")
	ns := 30000
	strWithUnit := "130s"
	strWithoutUnit := "130"

	fmt.Println(cast.ToDuration(d))               // 1m30s
	fmt.Println(cast.ToDuration(ns))              // 30µs
	fmt.Println(cast.ToDuration(strWithUnit))     // 2m10s
	fmt.Println(cast.ToDuration(strWithoutUnit))  // 130ns
}

