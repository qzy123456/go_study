
package main

import "fmt"

type Server struct {
	host string  // 主机地址
	weight int // 配置的权重
	currentWeight int  // 当前权重
}

func getSever(servers []*Server) (s *Server)  {

	allWeight := 0 // 总权重

	for _, server := range servers {
		if server == nil {
			return nil
		}
		// 每一轮选择都用自身的权重加到当前权重
		allWeight += server.weight
		server.currentWeight += server.weight

		// 当前未选中节点或当前节点比之前选中的节点权重高，那么更新当前选中的节点
		if s == nil || server.currentWeight > s.currentWeight{
			s = server
		}
	}

	s.currentWeight -= allWeight

	return
}


func main() {
	servers := []*Server{
		{"192.168.10.10", 5, 0},
		{"192.168.10.11", 2, 0},
		{"192.168.10.12", 1, 0},
	}

	for i := 0; i < 20; i++ {
		server := getSever(servers)

		if server == nil {
			continue
		}

		fmt.Printf("| %s | %d |\n", server.host, server.weight)
	}
}