package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"os"
)
var eg *xorm.EngineGroup
type UserInfo struct {
	Uid                int       `xorm:"uid"`
	Silver_coin        int       `xorm:"silver_coin"`
	Gold_coin          int       `xorm:"gold_coin"`
	Task_turnover      int       `xorm:"task_turnover"`
	Shop_grade         int       `xorm:"shop_grade"`
	Lives              int       `xorm:"lives"`
	Last               int       `xorm:"last"`
	New_level          int       `xorm:"new_level"`
	Stars              int       `xorm:"stars"`
	Chest_last         int       `xorm:"chest_last"`
	Chest_number       int       `xorm:"chest_number"`
}

type MyLogger struct {
}
type LoggerConfig struct {
	FileName            string `json:"filename"`
	Level               int    `json:"level"`    // 日志保存的时候的级别，默认是 Trace 级别
	Maxlines            int    `json:"maxlines"` // 每个文件保存的最大行数，默认值 1000000
	Maxsize             int    `json:"maxsize"`  // 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily               bool   `json:"daily"`    // 是否按照每天 logrotate，默认是 true
	Maxdays             int    `json:"maxdays"`  // 文件最多保存多少天，默认保存 7 天
	Rotate              bool   `json:"rotate"`   // 是否开启 logrotate，默认是 true
	Perm                string `json:"perm"`     // 日志文件权限
	RotatePerm          string `json:"rotateperm"`
	EnableFuncCallDepth bool   `json:"-"` // 输出文件名和行号
	LogFuncCallDepth    int    `json:"-"` // 函数调用层级
}
func main() {

//192.168.16.51:3306
	conns := []string{
		"root:root@tcp(192.168.20.185:3306)/ro_api_db2_dev?charset=utf8;",
		"root:root@tcp(192.168.16.51:3306)/ro_api_db2_dev?charset=utf8;",
		//"root:123456@tcp(192.168.16.78:3306)/ro_api_db2_dev?charset=utf8;",
		//"root:root@tcp(127.0.0.1:3306)/ro_api_db2_dev?charset=utf8;",
		//"mycat:123456@tcp(127.0.0.1:8066)/ro_api_db2_dev?charset=utf8;",
	}

	var err error
	eg, err = xorm.NewEngineGroup("mysql", conns)
	if err != nil{
		fmt.Println(err)
	}

	//config := rollingwriter.Config{
	//	LogPath:       "./logs",       //日志路径
	//	TimeTagFormat: "060102150405", //时间格式串
	//	FileName:      "mysql_exec",   //日志文件名
	//	MaxRemain:     3,              //配置日志最大存留数
	//	RollingPolicy:      rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
	//	RollingTimePattern: "* * * * * *",               //配置时间滚动策略
	//	RollingVolumeSize:  "1M",                        //配置截断文件下限大小
	//	WriterMode: "none",
	//	BufferWriterThershould: 256,
	//	// Compress will compress log file with gzip
	//	Compress: true,
	//}
	//
	//writer, err := rollingwriter.NewWriterFromConfig(&config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(writer)
	//
	//eg.SetLogger(logger)
	//eg.ShowSQL(true)
	//
	//fmt.Println(eg.Slave().DataSourceName())
	//eg.ShowSQL(true)
	user := new(UserInfo)
	////拼接条件
	////这个时候eg.Slave()竟然每次都是一样的,预计是编译之后，执行的都会是上次编译的代码把
	////has,_:= eg.Slave().Where("uid = ?", 57).Get(user)
	//has,_:= eg.Slave().Where("uid = ?", 57).Get(user)
	//if !has{
	//	fmt.Println("not found",err)
	//	return
	//}
	//eg.ShowSQL(true)
	//第二种简单模式
	f, err := os.OpenFile("sql.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		println(err.Error())
		return
	}
	eg.SetLogger(xorm.NewSimpleLogger(f))
	eg.ShowSQL(true)
	has,_:= eg.Slave().Where("uid = ?", 58).Get(user)
	if !has{
		fmt.Println("not found",err)

	}
	eg.ShowSQL(true)
	fmt.Println(has)
	fmt.Println(user)
}
func (UserInfo) TableName() string {

	return "user_info_1"
}
