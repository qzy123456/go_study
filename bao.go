package main

import "fmt"

/*
• 源⽂文件头部以 "package <name>" 声明包名称。
• 包由同⼀一⺫⽬目录下的多个源码⽂文件组成。
• 包名类似 namespace，与包所在⺫⽬目录名、编译⽂文件名⽆无关。 • ⺫⽬目录名最好不⽤用 main、all、std 这三个保留名称。
• 可执⾏行⽂文件必须包含 package main，⼊入⼝口函数 main。
说明:os.Args 返回命令⾏行参数，os.Exit 终⽌止进程。
要获取正确的可执⾏行⽂文件路径，可⽤用 filepath.Abs(exec.LookPath(os.Args[0]))。
包中成员以名称⾸首字⺟母⼤大⼩小写决定访问权限。
• public: ⾸首字⺟母⼤大写，可被包外访问。
• internal: ⾸首字⺟母⼩小写，仅包内成员可以访问。
该规则适⽤用于全局变量、全局常量、类型、结构字段、函数、⽅方法等
`
`导入包`
`
使⽤用包成员前，必须先⽤用 import 关键字导⼊入，但不能形成导⼊入循环。
import "相对⺫⽬目录/包主⽂文件名"
相对⺫⽬目录是指从 <workspace>/pkg/<os_arch> 开始的⼦子⺫⽬目录，以标准库为例:
在导⼊入时，可指定包成员访问⽅方式。⽐比如对包重命名，以避免同名冲突。
Go 学习笔记, 第 4 版
说明:os.Args 返回命令⾏行参数，os.Exit 终⽌止进程。 要获取正确的可执⾏行⽂文件路径，
可⽤用 filepath.Abs(exec.LookPath(os.Args[0]))。
import "fmt"      ->  /usr/local/go/pkg/darwin_amd64/fmt.a
import "os/exec"  ->  /usr/local/go/pkg/darwin_amd64/os/exec.a
import     "yuhen/test"// 默认模式: test.A
import  M  "yuhen/test"// 包重命名: M.A
import  .  "yuhen/test"// 简便模式: A
import  _  "yuhen/test"// ⾮非导⼊入模式: 仅让该包执⾏行初始化函数。
PS：
未使⽤用的导⼊入包，会被编译器视为错误 (不包括 "import _")。 ./main.go:4: imported and not used: "fmt"
*/
func main()  {
 fmt.Println(111)
}

