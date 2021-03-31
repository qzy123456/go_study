package main

import (
	"fmt"
	"net"
	"log"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"bufio"
	"os"
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
var listenAddr *net.UDPAddr
var socket *net.UDPConn

func main() {
	var err error

	// 设置log参数
	log.SetFlags(log.Lshortfile)

	// 用户集合
	users = make([]*net.UDPAddr, 0, 10)

	// 服务器地址
	serverAddr, err = net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	log.Println("serverAddr", err)

	port := 8000
PORT:

	// 本地地址
	listenAddr, err = net.ResolveUDPAddr("udp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
	}

	// 创建连接
	socket, err = net.ListenUDP("udp4", listenAddr)
	if err != nil {
		log.Println("连接失败!", err)
		port++
		goto PORT
		return
	}
	defer socket.Close()

	// 上线
	loginData := make([]byte, 0, 10)
	loginData = append(loginData, CMD_LOGIN)
	loginData = append(loginData, []byte("nickname")...)

	// 发送上线数据
	_, err = socket.WriteToUDP(loginData, serverAddr)
	if err != nil {
		log.Println("发送数据失败!", err)
		return
	}

	// 读取消息
	go readMsg()

	// 用户交互
	readCmd()
}

// 用户交互
func readCmd() {
	for {
		fmt.Printf("p2p > ")

		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			continue
		}

		var line = scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Println("read error: ", err)
			continue
		}

		switch {
		case strings.HasPrefix(line, "help"):
			fmt.Println("  list: show all user list\n  send: send message\n\tsend <id> <message>")
		case strings.HasPrefix(line, "list"):
			fmt.Println("user list:")
			for id, user := range users {
				fmt.Println(id+1, user.IP, user.Port)
			}
		case strings.HasPrefix(line, "send"):
			id := 0
			content := ""
			fmt.Sscanf(line, "send %d %s", &id, &content)

			if id <= 0 || id > len(users) {
				fmt.Printf("error: id %d not fund\n", id)
				continue
			}

			log.Printf("send message: %s %d, %s", users[id-1], id, content)

			sendData := make([]byte, 0, 100)
			sendData = append(sendData, CMD_MEG)
			sendData = append(sendData, []byte(content)...)

			n, err := socket.WriteToUDP(sendData, users[id-1])
			log.Println(n, err)
		default:
			fmt.Printf("command error: %s\nuse the 'help' command to get help\n", line)
		}
	}
}

func readMsg() {
	for {
		// 接收数据
		data := make([]byte, 1024)
		read, addr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Println("读取数据失败!", err)
			continue
		}
		log.Printf("UDP: %d, %s, %s\n", read, addr, hex.EncodeToString(data[:read]))

		switch data[0] {
		case CMD_LOGIN_RES:

		case CMD_LIST_RES:
			for i := 1; i < read; i+=6 {
				addrData := data[i:]
				touser := &net.UDPAddr{
					IP: net.IP(addrData[:4]),
					Port: int(binary.LittleEndian.Uint16(addrData[4:])),
				}

				coneData := make([]byte, 0, 10)
				coneData = append(coneData, CMD_CONE)
				coneData = append(coneData, []byte("nickname")...)

				socket.WriteToUDP(coneData, touser)
				log.Println("cone: ", touser, coneData)

				users = append(users, touser)
			}
		case CMD_PING:
			log.Printf("CMD_PING\n")
			pong_data := make([]byte, 0, 15)
			pong_data = append(pong_data, CMD_PONG, 1)
			n, err := socket.WriteTo(pong_data, addr)
			log.Println("CMD_PING: ", n, err)
		case CMD_PONG:

		case CMD_CONE:
			coneResData := make([]byte, 0, 10)
			coneResData = append(coneResData, CMD_CONE_RES)
			coneResData = append(coneResData, []byte("nickname")...)

			socket.WriteToUDP(coneResData, addr)
		case CMD_CONE_RES:
			log.Println("CMD_CONE_RES:", addr)
		case CMD_MEG:
			fmt.Println(string(data[1:read]))
		case CMD_MSG_RES:

		default:
			log.Println("default UDP: ", data[0])
		}
	}
}