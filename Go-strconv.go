package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseBool("1"))    // true
	fmt.Println(strconv.ParseBool("t"))    // true
	fmt.Println(strconv.ParseBool("T"))    // true
	fmt.Println(strconv.ParseBool("true")) // true
	fmt.Println(strconv.ParseBool("True")) // true
	fmt.Println(strconv.ParseBool("TRUE")) // true
	fmt.Println(strconv.ParseBool("TRue"))
	// false strconv.ParseBool: parsing "TRue": invalid syntax
	fmt.Println(strconv.ParseBool("0"))     // false
	fmt.Println(strconv.ParseBool("f"))     // false
	fmt.Println(strconv.ParseBool("F"))     // false
	fmt.Println(strconv.ParseBool("false")) // false
	fmt.Println(strconv.ParseBool("False")) // false
	fmt.Println(strconv.ParseBool("FALSE")) // false
	fmt.Println(strconv.ParseBool("FALse"))
	// false strconv.ParseBool: parsing "FAlse": invalid syntax
}