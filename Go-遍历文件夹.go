package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	//方式一
	filepath.Walk("./", func (path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	//方式二
	getFileLists("./")
}

func getFileLists(path string) {
	fs,_:= ioutil.ReadDir(path)
	for _,file:=range fs{
		if file.IsDir(){
			fmt.Println(path+file.Name())
			getFileLists(path+file.Name()+"/")
		}else{
			fmt.Println(path+file.Name())
		}
	}
}