package main

import (
	"flag"
	"fmt"
	"strings"
)

//定义一个类型，用于增加该类型方法
type sliceValue []string

//new一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

/*
Value接口：
type Value interface {
    String() string
    Set(string) error
}
实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
*/
func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

//flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's none of my business"
}

var (
	host   string
	dbName string
	port   int
	users  string
	pass   string
)
/*
可执行文件名 go run Go-flag.go  -slice="java,go"  最后将输出[java,go]
可执行文件名 go run Go-flag.go 最后将输出[default is me]
go run main.go -host=localhost -user=test -password=123456 -db_name=test -port=3306 -slice="java,go"
go run main.go -host=localhost noflag -user=test -password=123456 -db_name=test -port=3306
运行结果如下，可以看到解析-host参数之后遇到了noflag这样的非选项参数，flag就停止解析了，所以后面的参数都只输出了默认值。
数据库地址:localhost
数据库名称:
数据库用户:
数据库密码:
数据库端口:3306
 */
func main() {
	var languages []string
	//////////////////////////////////////////////////⚠️注意这个 slice 就是后面要跟的参数列表
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
	flag.StringVar(&host, "host", "", "数据库地址")
	flag.StringVar(&dbName, "db_name", "", "数据库名称")
	flag.StringVar(&users, "user", "", "数据库用户")
	flag.StringVar(&pass, "password", "", "数据库密码")
	flag.IntVar(&port, "port", 3306, "数据库端口")
	flag.Parse()
	fmt.Printf("数据库地址:%s\n", host)
	fmt.Printf("数据库名称:%s\n", dbName)
	fmt.Printf("数据库用户:%s\n", users)
	fmt.Printf("数据库密码:%s\n", pass)
	fmt.Printf("数据库端口:%d\n", port)
	//打印结果slice接收到的值
	fmt.Println(languages)
	//int和uint互相转换
	var n int = -100
	var un uint = uint(n)
	fmt.Println(un)
	// //自己创建一个命令行参数的集合
	//    var flagSet = flag.NewFlagSet("my flag", flag.ExitOnError)
	//
	//    host := flagSet.String("host", "", "数据库地址")
	//    dbName := flagSet.String("db_name", "", "数据库名称")
	//    user := flagSet.String("user", "", "数据库用户")
	//    password := flagSet.String("password", "", "数据库密码")
	//    port := flagSet.Int("port", 3306, "数据库端口")
	//
	//    //解析命令行参数，从os.Args的第二个元素开始，第一个元素是命令本身
	//    flagSet.Parse(os.Args[1:])
	//
	//    fmt.Printf("数据库地址:%s\n", *host)
	//    fmt.Printf("数据库名称:%s\n", *dbName)
	//    fmt.Printf("数据库用户:%s\n", *user)
	//    fmt.Printf("数据库密码:%s\n", *password)
	//    fmt.Printf("数据库端口:%d\n", *port)
	// //模拟os.Args数组，定义一个参数数组
	//    var params = []string{"-host", "127.0.0.1", "-db_name", "test", "-user", "test", "-password", "abcdef", "-port", "13306"}
	//
	//    var flagSet = flag.NewFlagSet("my flag", flag.ExitOnError)
	//
	//    host := flagSet.String("host", "", "数据库地址")
	//    dbName := flagSet.String("db_name", "", "数据库名称")
	//    user := flagSet.String("user", "", "数据库用户")
	//    password := flagSet.String("password", "", "数据库密码")
	//    port := flagSet.Int("port", 3306, "数据库端口")
	//
	//    //解析自定义的参数数组
	//    flagSet.Parse(params)
	//    fmt.Printf("数据库地址:%s\n", *host)
	//    fmt.Printf("数据库名称:%s\n", *dbName)
	//    fmt.Printf("数据库用户:%s\n", *user)
	//    fmt.Printf("数据库密码:%s\n", *password)
	//    fmt.Printf("数据库端口:%d\n", *port)
	//运行程序，在命令后面不需要跟命令行参数，如下：
	//go run main.go
	//我们在使用Linux命令的时候，发现很多命令的参数是有分短选项和长选项的，不过flag库并不支持短选项；当然也有变通的方式，
	//比如我们可以自己定义一个长选项和短选项，
	//var port int
	//flag.IntVar(&port, "p", 3306, "数据库端口")
	//flag.IntVar(&port, "port", 3306, "数据库端口")
	//flag.Parse()
	//fmt.Println(port)
	//上面的程序中，我们定义了p和port两个参数，并将其绑定到变量port，因此通过下面两条命令都可以获取参数：
	//go run main.go -p 111
	//go run main.go -port 111
}
