package main

import (
	"fmt"
)


type Requestt struct {
	Method string
	Body   string
}

type RequesttOption func() (*Requestt, error)

func (r *Requestt) WithCreate() RequesttOption {
	return func() (request *Requestt, e error) {
		return r, nil
	}
}

func (opt RequesttOption) WithMethod(method string) RequesttOption {
	return func() (request *Requestt, e error) {
		r, err := opt()
		if err != nil {
			return r, err
		}
		r.Method = method
		return r, err
	}
}

func (opt RequesttOption) WithBody(body string) RequesttOption {
	return func() (request *Requestt, e error) {
		r, err := opt()
		if err != nil {
			return r, err
		}
		r.Body = body
		return r, err
	}
}

func (opt RequesttOption) Send() error {
	r, err := opt()
	if err != nil {
		return err
	}
	fmt.Printf("发送%s请求,请求体为%s", r.Method, r.Body)
	return nil
}
func main()  {
	req := new(Requestt)
	//err := req.WithCreate().WithMethod("GET").WithBody("获取图片").Send()  //正常调用
	err := req.WithCreate().WithBody("获取图片").WithMethod("GET").Send()
	if err != nil {
		panic(err)
	}
	//截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组⻓度为 l。在操作符 [i:j]
	//中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的⻓度，
	// 截取得到的切⽚⻓度和容ᰁ计算⽅法是 ji、l-i。操作符 [i:j:k]，k 主要是⽤来限制切⽚的容量，但是不能⼤于数组的⻓度 l，截取得到的切⽚⻓度
	//和容ᰁ计算⽅法是 j-i、k-i。
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t,cap(t),len(t))
	//4 1 1
}

