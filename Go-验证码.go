package main

import (
    "image/color"
    "image/png"
    "net/http"

    "github.com/afocus/captcha"
)

var capp *captcha.Captcha

func main() {
    //创建1个句柄
    capp = captcha.New()
    //通过句柄调用 字体文件
     if err := capp.SetFont("/Users/artist/go/src/github.com/afocus/captcha/examples/comic.ttf"); err != nil {
        panic(err.Error())
    }
    //设置图片的大小
    capp.SetSize(128, 64)
    // 设置干扰强度
    capp.SetDisturbance(captcha.MEDIUM)
    // 设置前景色 可以多个 随机替换文字颜色 默认黑色
    capp.SetFrontColor(color.RGBA{255, 255, 255, 255})
    // 设置背景色 可以多个 随机替换背景色 默认白色
    capp.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

    http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
        //参数一：验证码的个数 参数二：验证码内容组成
        //生成图片 返回图片和 字符串(图片内容的文本形式)
        img, str := capp.Create(4, captcha.NUM)
        png.Encode(w, img)
        println(str)
    })

    http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
        str := r.URL.RawQuery
        img := capp.CreateCustom(str)
        png.Encode(w, img)
    })

    http.ListenAndServe(":8085", nil)

}
