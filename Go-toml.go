package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"reflect"
)

type database struct {
	Server string
	Ports  []int
}

type config struct {
	Database database
	MultiDbTable map[string]map[string]int64
	DbWeight   map[string]int64
}
func main() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed,err:", err)
	}
	var c config
	viper.Unmarshal(&c)
	fmt.Println(c.Database.Ports[1])
	fmt.Println(c.MultiDbTable)
	fmt.Println(c.DbWeight)

}

func Struct2map(obj interface{}) (data map[string]interface{}, err error) {
	// 通过反射将结构体转换成map
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		fileName, ok := objT.Field(i).Tag.Lookup("json")
		if ok {
			data[fileName] = objV.Field(i).Interface()
		}else{
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
	}
	return data, nil
}
