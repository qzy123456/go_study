package main
import (
	"fmt"
	"time"
)

func main()  {
	//创建一个有缓存的channel，容量为3
	ch := make(chan int, 3)
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {    //改成i<10试试
			ch <- i //不会阻塞，ch容量为3
			fmt.Printf("子协程[%d]：len(ch)=%d,cap(ch)=%d\n", i, len(ch), cap(ch))
		}
	}()

	//延时
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {    //改成i<10试试(报错，因为容量超过最大值了)
		num := <-ch //读取管道中内容，没有内容前，阻塞
		fmt.Println("num =", num)
	}



	//创建一个定时器，设置时间为2s，2s后往time通道写内容（当前时间）
	timer := time.NewTimer(2*time.Second)
	fmt.Println("Current time :",time.Now())

	// 2s后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C    //channel没有数据前后阻塞
	fmt.Println("t = ",t)

	var tt int64 = time.Now().Unix()
	var s string = time.Unix(tt, 0).Format("2006-01-02 15:04:05")
	t1, _ := time.Parse( "2006-01-02 15:04:05",s)
	datetime_str_to_timestamp := t1.Unix()
	println(s)
	println(datetime_str_to_timestamp)



}