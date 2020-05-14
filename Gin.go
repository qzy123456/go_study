package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 这个处理器可以匹配 /user/john ， 但是它不会匹配 /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个可以匹配 /user/john 和 /user/john/send
	// 如果没有其他的路由匹配 /user/john ， 它将重定向到 /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// 查询字符串参数使用现有的底层 request 对象解析。
	// 请求响应匹配的 URL： /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// 这个是 c.Request.URL.Query().Get("lastname") 的快捷方式。
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	// 简单组： v1
	//v1 := r.Group("/v1")
	//{
	//	v1.POST("/login", loginEndpoint) //loginEndpoint到时候
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}
	//在不同的域中使用 JSONP 从一个服务器请求数据。如果请求参数中存在 callback，添加 callback 到 response body
	r.GET("/JSONP?callback=x", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback 是 x
		// 将会输出  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})
	//使用 SecureJSON 来防止 json 劫持。如果给定的结构体是数组值，默认预置 "while(1)," 到 response body 。
	// 你也可以使用自己的安装 json 前缀
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将会输出  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
	//使用 AsciiJSON 生成仅有 ASCII 字符的 JSON，非 ASCII 字符将会被转义 。
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 将会输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
    //从reader提供数据
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
   //使用 LoadHTMLGlob () 或 LoadHTMLFiles ()
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	//templates/index.tmpl
	// <html>
	//    <h1>
	//        {{ .title }}
	//    </h1>
	//</html>
	//自定义中间件
	r.Use(Logger())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// 它将打印： "12345"
		log.Println(example)
	})

	//使用自定义结构绑定表单数据请求
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)
	// curl "http://localhost:8899/getb?field_a=hello&field_b=world"
	//{"a":{"FieldA":"hello"},"b":"world"}
	//$ curl "http://localhost:8899/getc?field_a=hello&field_c=world"
	//{"a":{"FieldA":"hello"},"c":"world"}
	//$ curl "http://localhost:8899/getd?field_x=hello&field_d=world"
	//{"d":"world","x":{"FieldX":"hello"}}
	_ = r.Run(":8899") // listen and serve on 0.0.0.0:8080
}
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置简单的变量
		c.Set("example", "12345")

		// 在请求之前

		c.Next()

		// 在请求之后
		latency := time.Since(t)
		log.Print(latency)

		// 记录我们的访问状态
		status := c.Writer.Status()
		log.Println(status)
	}
}