package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"sync"
	"time"
)

var mu sync.RWMutex

func help() string {
	return (`
    please choose options:
        - register : 注册。            格式:"register"
        - login    : 登录。            格式:"login"
        - logoff   : 注销登录。         格式:"logoff"
        - exit     : 退出系统。         格式:"exit"
        - add      : 添加好友。         格式:"add"
        - delete   : 删除好友。         格式:"delete"
        - list     : 查看好友列表。      格式:"list"
        - sendTo   : 给某位好友发送消息。 格式:"sendTo"
        - sendAll  : 给全部好友发送消息。 格式:"sendAll"
        `)
}

/*
Clients is a demo
*/
type Clients struct {
	userName string
}
type userCH chan<- string

var (
	messages = make(chan string)
	entering = make(chan userCH)
	leaving  = make(chan userCH)
	clients  = make(map[string]userCH)
	db       *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1)/test1?charset=utf8")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//db.Query("create database if not exists chat;")
	db.Query("create table if not exists user(id int auto_increment primary key, userName varchar(20) unique, password varchar(20) not null, tel varchar(20) unique);")
	//创建对应的表
	createInfo()
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//启动一个携程  专门去跑客户端传来的链接
		go handles(conn)
	}
}
//携程～～用来处理客户端传来的链接～～～
func handles(conn net.Conn) {
	var client *Clients
	ch := make(chan string)
	//写携程～～用来专门反馈给客户端消息
	go clientWriters(conn, ch)
	//地址！！类似127。0。0。1：6330 这种
	who := conn.RemoteAddr().String()
	//刚连接上的欢迎语～～以及让其选择登陆，注册～～
	ch <- "welcome! " + who + "\n"
	ch <- help() + "\n"
	ch <- ">>"
	//读取输入框
	input := bufio.NewScanner(conn)
	for input.Scan() {
		var op = input.Text()
		fmt.Println(op)
		switch op {
		case "register":
			//register
			ok := registers(input)
			if !ok {
				ch <- "Fail register=_=||\nmaybe your userName or phoneNumber is invalid\n"
			} else {
				ch <- "Success register!\n"
			}
		case "login":
			//login
			//TODO 多重登陆问题，做个强制下线？
			client = logins(input)
			if client != nil {
				clients[client.userName] = ch
				ch <- "Success login!\n" + "Your new messages:\n"
				client.toRead(ch)
			} else {
				ch <- "Fail login=_=||maybe your userName or password is wrong,please check them carefully:)\n"
			}
		//下线（清空连接）～～非exit（退出程序）～～
		case "logoff":
			if client == nil {
				ch <- "Please login first:)\n"
			} else {
				ch <- "you logoff successfully!\n"
				ch <- help() + "\n"
				clients[client.userName] = nil
				client = nil
			}
		case "add":
			var userName string
			if input.Scan() {
				userName = input.Text()
			}
			if client == nil {
				ch <- "Please login first:)\n"
			} else {
				ok := client.addFriend(userName)
				if ok {
					ch <- "add successfully!\n"
				} else {
					ch <- "maybe no such user=_=||\n"
				}
			}
		case "delete":
			var userName string
			if input.Scan() {
				userName = input.Text()
			}
			if client == nil {
				ch <- "Please login first:)\n"
			} else {
				ok := client.deleteFriend(userName)
				if ok {
					ch <- "delete successfully!\n"
				} else {
					ch <- "maybe no such user=_=||\n"
				}
			}
		case "list":
			if client == nil {
				ch <- "Please login first:)\n"
			} else {
				ch <- client.list() + "\n"
			}
		case "sendTo":
			if client == nil {
				ch <- "Please login first:)\n"
			} else {
				var userName, content string
				if input.Scan() {
					userName = input.Text()
				}
				if input.Scan() {
					content = input.Text()
				}
				ok := client.sendTo(userName, content)
				if ok {
					ch <- "发送成功\n"
				} else {
					ch <- "发送失败，可能你没有这个好友:)\n"
				}
			}
		case "sendAll":
			if client == nil {
				ch <- "Please login first:)\n"
			} else {
				var content string
				if input.Scan() {
					content = input.Text()
				}
				client.sendAll(content)
			}
		default:
			ch <- "无法识别的命令=_=||请重新输入正确的命令:)\n"
			ch <- help() + "\n"
		}
		ch <- ">>"
	}

	conn.Close()
}

func clientWriters(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprint(conn, msg)
	}
}

//读取user，判断手机号和用户名是否重复注册
//将注册信息写入user
//如果成功就创建messages、logs、friends表
func registers(cin *bufio.Scanner) bool {
	var userName, password, tel string
	if cin.Scan() {
		userName = cin.Text()
	}
	if cin.Scan() {
		password = cin.Text()
	}
	if cin.Scan() {
		tel = cin.Text()
	}
	mu.Lock()
	defer mu.Unlock()
	stmt, err := db.Prepare("insert user set userName=?,password=?,tel=?")
	if err != nil {
		log.Fatalln("in func register:insert table user")
	}
	res, err := stmt.Exec(userName, password, tel)
	if err != nil {
		fmt.Println("in func register:insert table user went wrong!!")
		return false //插入失败，可能是用户名或者手机号已被注册
	}
	ok, _ := res.RowsAffected()
	if ok != 1 { //其实这里应该不会有问题了
		fmt.Println("in func register:insert table user went wrong!!")
		return false //插入失败，可能是用户名或者手机号已被注册
	}
	fmt.Printf("%s successfully register!\n", userName)
	//创建表
	return true
}

//创建表
func createInfo() {
	stmt, _ := db.Prepare("create table friends (id int primary key auto_increment, userName varchar(20) unique)")
	stmt.Exec()
	stmt, _ = db.Prepare("create table logs (id int primary key auto_increment, wtime datetime, fromwho varchar(20), towho varchar(20), msg text)")
	stmt.Exec()
	stmt, _ = db.Prepare("create table messages (id int primary key auto_increment, wtime datetime, fromwho varchar(20), msg text)")
	stmt.Exec()
}

//读取user,判断是否有这个用户以及密码是否正确
func logins(cin *bufio.Scanner) *Clients {
	var userName, password string
	if cin.Scan() {
		userName = cin.Text()
	}
	if cin.Scan() {
		password = cin.Text()
	}
	mu.RLock()
	defer mu.RUnlock()
	rows, err := db.Query("select * from user where userName=? and password=?", userName, password)
	if err != nil {
		log.Fatalln(err)
	}
	if rows.Next() {
		var client Clients
		client.userName = userName
		return &client
	}
	rows.Close()
	return nil
}
//读取离线消息
func (c *Clients) toRead(ch chan string) {
	mu.Lock() //写日志加锁
	defer mu.Unlock()
    //从messgae表中读取消息
	rows, err := db.Query("select wtime, fromwho, msg from messages")
	defer rows.Close()
	if err != nil {
		log.Fatalln("in func toRead: ", err)
	}
      //把离线表的消息 插入到消息表中去
	stmt, err := db.Prepare("insert logs set wtime = ?, fromwho = ?, towho = ?, msg = ?")
	for rows.Next() {
		var t, who, msg string
		err := rows.Scan(&t, &who, &msg)
		if err != nil {
			log.Fatalln("in func toRead:", err)
		}
		//拼接消息字段
		temp := t + " " + who + " " + msg
		ch <- temp + "\n"
		//执行插入
		_, err = stmt.Exec(t, who, c.userName, msg)
		if err != nil {
			log.Fatalln("fail to insert into logs", err)
		}
	}
	//最后把离线表中的消息给清空
	stmt, _ = db.Prepare("delete from  messages" )
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalln(err)
	}
}
//添加好友～～默认是添加成功（好友表  有毛病～～只有好友名字，没有对应的自己的名字～需要修改把）
func (c *Clients) addFriend(userName string) bool {
	mu.RLock()
	rows, _ := db.Query("select userName from user where userName = ?", userName)
	mu.RUnlock()
	defer rows.Close()
	if rows.Next() {
		//这个应该不用锁，只有自己读写
		stmt, err := db.Prepare("insert friends set userName = ?")
		if err != nil {
			log.Fatalln(err)
		}
		_, err = stmt.Exec(userName)
		if err != nil {
			log.Fatalln(err)
		}
		return true
	}
	return false
}
//删除好友
func (c *Clients) deleteFriend(userName string) bool {
	rows, err := db.Query("select userName from friends where userName = ?", userName)
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if rows.Next() {
		stmt, _ := db.Prepare("delete from friends where userName = ?")
		_, err := stmt.Exec(userName)
		if err != nil {
			log.Fatalln(err)
		}
		return true
	}
	return false

}
//好友列表
func (c *Clients) list() string {
	res := "Your friends list:"
	rows, _ := db.Query("select userName from friends ")
	var userName string
	for rows.Next() {
		rows.Scan(&userName)
		res += "\n\t" + userName
	}
	return res
}
//发送给固定的好友
func (c *Clients) sendTo(userName, content string) bool {
	mu.Lock()
	defer mu.Unlock()
	//从好友表中查出好友是否存在
	err := db.QueryRow("select userName from friends where userName = ?", userName).Scan(&userName)
	if err != nil {
		return false
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	//拼接发送的信息格式
	msg := "\n" + t + " " + c.userName + " " + content
	//当前好友在线
	if clients[userName] != nil {
		clients[userName] <- msg + "\n"
		clients[userName] <- ">>"
		//存储发送的消息
		stmt, err := db.Prepare("insert logs set wtime=?, fromwho=?, towho=?, msg=?")
		_, err = stmt.Exec(t, c.userName, userName, content)
		if err != nil {
			log.Fatalln("in func sendTo: table(logsOf"+userName+"): ", err)
		}
     //离线要把消息存储到离线表中，下次用户上线，再查询离线消息  发送给他
	} else {
		fmt.Println("这时候掉线了")
		stmt, err := db.Prepare("insert messages set wtime=?, fromwho=?, msg=?")
		_, err = stmt.Exec(t, c.userName, content)
		if err != nil {
			log.Fatalln("in func sendTo: table(messagesOf"+userName+"): ", err)
		}
	}
	return true
}
//发送给所有人
func (c *Clients) sendAll(content string) bool {
	//枷锁
	mu.Lock()
	defer mu.Unlock()
	//时间
	t := time.Now().Format("2006-01-02 15:04:05")
	//拼接msg
	msg := "\n" + t + " " + c.userName + " " + content
	//查出所有的好友～～这时也包含自己
	rows, _ := db.Query("select userName from friends")
	defer rows.Close()
	var userName string
	// 一次一次的循环查找
	for rows.Next() {
		rows.Scan(&userName)
		//当前用户在线
		if clients[userName] != nil {
			//发消息
			clients[userName] <- msg + "\n"
			clients[userName] <- ">>"
			//把发送的消息信息存储到日志表中
			stmt, err := db.Prepare("insert logs set wtime=?, fromwho=?, towho=?, msg=?")
			_, err = stmt.Exec(t, c.userName, userName, content)
			if err != nil {
				log.Fatalln("in func sendAll: insert logsOf: ", err)
			}
		//当前用户不在线，那么就要先把信息存储到离线表中,下次用户上线直接获取离线信息
		} else {
			stmt, err := db.Prepare("insert messages set wtime=?, fromwho=?, msg=?")
			_, err = stmt.Exec(t, c.userName, userName, content)
			if err != nil {
				log.Fatalln("in func sendAll: insert messagesOf: ", err)
			}
		}
	}
	return true
}
