package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpPost() error {

	client := http.Client{}

	data := "loginId=1498321166%40126.com&password=aaa123&exp=-1" //设置提交数据

	request, err := http.NewRequest("POST", "http://www.iwordnet.com/registerInner.htm", strings.NewReader(data)) //请求

	if err != nil {

		return err // handle error

	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8") //设置Content-Type

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36") //设置User-Agent

	response, err := client.Do(request) //返回

	if err != nil {

		return err

	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {

		return err
	}
	fmt.Println(string(body)) //打印返回文本
	//{"id":6280979,"message":"success","t":"P3230GK0Iv8YRnSJXkYYVBA380KDfpdDqg4IGJjJzipaKA1CE9","n":"ZM_78K3jfl6Uf","ctime":1531913582000,"success":true}
	return nil
}


func main() {

	httpPost() //调用函数

}