package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var Slice  = []int {1, 5, 10,}
func GetData1() {
	//设置超时
	client := &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: time.Second * 2,
		},
		Timeout:time.Second * 2,
	}
	resp, err := client.Get("http://api.map.baidu.com/place/v2/suggestion?query=广州市天河区正佳广场&region=广州&city_limit=true&output=json&ak=yX8nC9Qzpckek7lY9gGWmlD4TFcA2tzYx3")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
func httpPost1() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main() {
	ss := &http.Server{
		Addr:":1212",
		Handler:nil,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	  if err :=ss.ListenAndServe(); err!=nil{
		println(err)
	  }
	GetData1()
	 httpPost1()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}
	wg.Wait()

	if true {
		defer fmt.Println(1)
	} else {
		defer fmt.Println(2)
	}
	fmt.Println(3)

	x := 1

	{

		x := 2

		fmt.Print(x)

	}

	fmt.Println(x) //21

	strs := []string{"one", "two", "three"}

	for _, s := range strs {

		go func() {

			fmt.Printf("%s ", s)

		}()

	}
  //three three three

}

