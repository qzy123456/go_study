package main

import (
	"fmt"
	"github.com/oxequa/grace"
)

func example() (e error){
	defer grace.Recover(&e) // save recover error and stack trace to e
	numbers := []int{1, 2}
	fmt.Println(numbers[3]) // panic out of index
	return
}

func main() {
	err := example() // no panic occur
	fmt.Println(err)
	fmt.Println("End")
}
