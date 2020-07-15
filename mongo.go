package main

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	_ "gopkg.in/mgo.v2/bson"
	_ "log"
)
type Person struct {
	NAME  string
	PHONE string
	CONTENT string
}
type Men1 struct {
	Persons []Person
}

const (
	URL = "mongodb://root:root@127.0.0.1:27017" //连接mongoDB启动服务的端口号 你得先启动mongoDB服务
	URL1 = "mongodb://127.0.0.1:27017" //连接mongoDB启动服务的端口号 你得先启动mongoDB服务
)

func main() {

	session, err := mgo.Dial(URL) //连接数据库
	//连接失败
	if err != nil {
		panic(err)
	}
	defer session.Close()
	//Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("test") //数据库名称

	collection := db.C("person") //如果该集合已经存在的话，则直接返回

	//*****集合中元素数目********
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)

	//*******插入元素*******
	//temp := &Person{
	//	PHONE: "7017986",
	//	NAME:  "Ale",
	//}
	var allLog = make(map[string]interface{})
	//基础参数
	baseLog := map[string]string{
		"event_uuid":"fdasdasd",

	}
	//拼接key，每个用户的key是不一样的
	allLog["baseLog"] = baseLog
	basedatas,_ := json.Marshal(allLog)
	//str := `{"servers":[{"serverName":"TianJin","serverIP":"127.0.0.1"},
//{"serverName":"Beijing","serverIP":"127.0.0.2"}]}`
	fmt.Println("xxxxx ", basedatas)
	//一次可以插入多个对象 插入两个Person对象
	//err = collection.Insert(&Person{"Ale", "13798245114",str}, temp)
	//if err != nil {
	//	panic(err)
   //	}

	//*****查询单条数据*******
	result := Person{}

	err = collection.Find(bson.M{"phone": "123456"}).One(&result) //查询单条phone为13798245114的结果
	fmt.Println("Phone:", result.NAME, result.PHONE)                   //输出单条phone为13798245114的结果

	//*****查询多条数据*******
	//var personAll Men1 //存放结果
	countryCapitalMap := make(map[int]map[string]string)
	subMapB := make(map[string]string)
	iter := collection.Find(nil).Iter()
	var ii int = 0;
	//一次一次的查
	for iter.Next(&result) {

		subMapB["name"] = result.NAME
		subMapB["phone"] = result.PHONE
		if result.CONTENT != ""{
			subMapB["content"] = result.CONTENT
		}
         ii++
		countryCapitalMap[ii] = subMapB
	}
	fmt.Println("===",countryCapitalMap)
	//var users []Person
	//err = collection.Find(nil).All(&users)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Printf("personAll: %v\n", users)
	//for i:= 0;i<len(users);i++{
	//	fmt.Print(users[i].NAME,"\t")
	//}

	//*******更新数据**********
	//修改所有name为ddd的对象成name为ddd
	_, err = collection.UpdateAll(bson.M{"name": "Ale"}, bson.M{"$set": bson.M{"name": "ddd"}})

	//修改name为ddd的对象成phone为666666
	err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "666666"}})

	//修改所有name为ddd的对象成name为xiaomin,phone为123456
	_, err = collection.UpdateAll(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"name": "xiaomin", "phone": "123456"}})

	//******删除所有name为xiaomin的数据************
	//_, err = collection.RemoveAll(bson.M{"name": "xiaomin"})
}
