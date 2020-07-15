package main

import (
	"fmt"
	"net/http"
)

func indexHandler( w http.ResponseWriter , r *http.Request)  {
	fmt.Fprintf(w,"hello word")
}

func htmlHadler( w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","text/html")

	html :=`
<!doctype html>    
<META http-equiv="Content-Type" content="text/html" charset="utf-8">    
<html lang="zh-CN">            
<head>                    
<title>Golang</title> 
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0;" />            
</head>           
 <body>             
   <div id="app">Welcome!</div>           
 </body>    
</html>
`
	fmt.Fprintf(w,html)
}

type myHandler struct {}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This message is from myHandler."))
}

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/",http.HandlerFunc(indexHandler))
	mux.HandleFunc("/welcome",htmlHadler)
	mux.Handle("/wel",&myHandler{})
	http.ListenAndServe(":8910",mux)
}

