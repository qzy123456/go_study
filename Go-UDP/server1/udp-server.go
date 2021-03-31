package main

import (
	"net"
	"encoding/binary"
	"encoding/hex"
	"log"
	"time"
)

const (
	// login client -> server
	CMD_LOGIN byte = byte(iota)
	CMD_LOGIN_RES

	// user list server -> client
	CMD_LIST
	CMD_LIST_RES

	// ping server -> client
	CMD_PING
	CMD_PONG

	// cone client -> client
	CMD_CONE
	CMD_CONE_RES

	// message client -> client
	CMD_MEG
	CMD_MSG_RES
)

var users []*net.UDPAddr
var serverAddr *net.UDPAddr
var socket *net.UDPConn

func main() {
	var err error

	// 设置log参数
	log.SetFlags(log.Lshortfile)

	// 用户集合
	users = make([]*net.UDPAddr, 0, 10)

	// 服务器地址
	serverAddr, err = net.ResolveUDPAddr("udp4", ":8080")
	log.Println(err)

	// 创建监听
	socket, err = net.ListenUDP("udp4", serverAddr)
	if err != nil {
		log.Println("监听失败! ", err)
		return
	}
	defer socket.Close()

	// ping
	 //go ping()

	// 处理
	for {
		// 处理数据报文
		data := make([]byte, 4096)
		read, addr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Println("读取数据失败!", err)
			continue
		}
		log.Printf("UDP: %d, %s, %s\n", read, addr, hex.EncodeToString(data[:read]))

		switch data[0] {
		case CMD_LOGIN:
			// 新上线用户的信息
			touser_list_data := make([]byte, 0, 15)
			touser_list_data = append(touser_list_data, CMD_LIST_RES, 0, 0, 0, 0, 0, 0)
			// touser_list_data = append(touser_list_data, addr.IP...)
			copy(touser_list_data[1:5], addr.IP)
			binary.LittleEndian.PutUint16(touser_list_data[5:], uint16(addr.Port))

			log.Println("touser_list_data:", hex.EncodeToString(touser_list_data))

			// 构建在线的用户信息列表
			userListData := make([]byte, 0, 100)
			userListData = append(userListData, CMD_LIST_RES)

			// 通知所有在线用户,有新用户上线
			for _, touser := range users {
				// 添加在线用户信息到列表
				// user_list_data = append(user_list_data, touser.IP...)
				userListData = append(userListData, 0, 0, 0, 0, 0, 0)
				copy(userListData[len(userListData)-6:], touser.IP)
				binary.LittleEndian.PutUint16(userListData[len(userListData)-2:], uint16(touser.Port))

				// 给在线用户发送数据
				socket.WriteToUDP(touser_list_data, touser)
			}

			// 给新上限用户发送在线用户的列表
			log.Println("user_list_data:", hex.EncodeToString(userListData))
			socket.WriteToUDP(userListData, addr)

			// 将新用户存储
			users = append(users, addr)
		case CMD_LOGIN_RES:
		case CMD_LIST:
		case CMD_LIST_RES:

		case CMD_PING:
		case CMD_PONG:
			log.Println("CMD_PONG udp: ", addr)
		case CMD_CONE:
		case CMD_CONE_RES:
		case CMD_MEG:
		case CMD_MSG_RES:
		default:
			log.Println("default udp: ", addr)
		}
	}
}

func ping() {
	pingData := make([]byte, 0, 15)
	pingData = append(pingData, CMD_PING, 0)

	for {
		for _, touser := range users {
			socket.WriteToUDP(pingData, touser)
		}

		time.Sleep(5 * time.Second)
	}
}
