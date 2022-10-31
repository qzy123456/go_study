package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/*
 * Golang下针对特定文件监控 以及热更新
 */

type HotLoad struct {
}

func New() *HotLoad {
	return &HotLoad{}
}

// 更新配置方法
type LoadFunc func(filePath string)

func (h *HotLoad) Load(filePath string, dstFunc LoadFunc) {
	md51 := ""
	for {
		select {
		case <-time.After(time.Millisecond * 200):
			md5, e := md5File(filePath)
			if e != nil {
				log.Fatalln(e)
			}
			if md51 != md5 {
				// 如果文集hash发生改变
				// 重新读取文件并调用dstFunc方法 进行热更新
				dstFunc(filePath)
				md51 = md5
			}
			continue
		}
	}
}

// 定义更新配置方法
func update(filePath string) {
	log.Println("更新了配置")
	bytes, e := ioutil.ReadFile(filePath)
	if e != nil {
		log.Fatalln(e)
	}

	log.Println(string(bytes))
}

// 计算指定文件的hash值
func md5File(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	load := New()           // 初始化hotload
	path := "my.ini"
	load.Load(path, update) // 配置监控
}