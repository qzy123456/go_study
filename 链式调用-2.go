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
}

