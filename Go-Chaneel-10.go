package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var printChar chan int
func prinNums() {
	defer wg.Done()
	for i:=0;i < 2; i++ {
		printChar <- 1111
		fmt.Println(<-printChar)
	}
}

func printChars(){
	defer wg.Done()
	for i:=0;i < 2; i++ {
		fmt.Println("阻1")
		fmt.Println(<-printChar)
		fmt.Println("阻2")
		fmt.Println("出来1")
		printChar <- 1222
		fmt.Println("出来2")
	}
}
func main(){
	printChar = make(chan int)

	wg.Add(2)

	go prinNums()
	go printChars()

	wg.Wait()
}


