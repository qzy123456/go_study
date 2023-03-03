package main

import (
	"fmt"
	"time"
)

type Server1 struct {
	appr     string
	port     string
	timout   time.Duration
	maxConns int
}
type Options func(*Server1)

func Timeout(duration time.Duration) Options {
	return func(server1 *Server1) {
		server1.timout = duration
	}
}

func Maxconns(maxcons int) Options {
	return func(server1 *Server1) {
		server1.maxConns = maxcons
	}
}

func NewServer1(add, port string, options ...Options) (*Server1, error) {
	s := &Server1{
		appr:     add,
		port:     port,
		timout:   time.Second * 10,
		maxConns: 20,
	}
	for _, option := range options {
		option(s)
	}
	return s, nil
}

func main() {
	server, err := NewServer1("127.0.0.1", "8888", Timeout(10), Maxconns(20))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(server)
}
