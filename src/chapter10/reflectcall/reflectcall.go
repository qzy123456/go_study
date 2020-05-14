package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"reflect"
)

// 普通函数
func add(a, b int) int {

	return a + b
}
var (fil  string
n int64
err error
)
func main() {

	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)

	// 构造函数参数，传入两个整形值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

	// 反射调用函数
	retList := funcValue.Call(paramList)

	// 获取第一个返回值，取整数值
	fmt.Println(retList[0].Int())
	fil ,n ,err = fetch("http://www.baidu.com")
    fmt.Println(fil,n,err)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	fmt.Println(resp.Request.URL)
	return
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
