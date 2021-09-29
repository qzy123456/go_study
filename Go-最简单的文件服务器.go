package main
import (
	"net/http"
)
func main(){
	//最简单的文件服务器
	h := http.FileServer(http.Dir("."))
	http.ListenAndServe(":8001",  h)
}