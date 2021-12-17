package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

func main() {
	_ := iris.New()
	//创建websocket服务器
	ws := websocket.New(websocket.Config{

	})
	ws.OnConnection(func(c websocket.Connection) {
		c.Join("/push")
		go func() {
			for {
				//（这里就是遇到的 坑当时一直钻牛角尖怎么把这里从man函数拿出去，其实只要直接勇哥chananel传进来就行，推商铺那个成功之后在一个chananel回去就可以了，）
				stringdata := "111"
				conns := ws.GetConnectionsByRoom("/push")
				for _, v := range conns { //遍历所有的链接,发送（这里要遍历所有的链接，不然第一次推送会推送给第一个链接上的，第二次推送就会推送给第二个链接上的，不是所有的的都会推送）
					err := v.EmitMessage([]byte(stringdata))
					if err != nil {
						v.Disconnect()
					}
				}
			}
		}()
	})


}