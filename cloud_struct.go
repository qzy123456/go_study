package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Apps struct {
	Sys    Sys    `json:"sys"`
	Device Device `json:"device"`
	App1    App1    `json:"app"`
	Data2   string   `json:"data"`
}
type Sys struct {
	AppID     string `json:"appId"`
	MyCompany string `json:"MyCompany"`
	MyGame    string `json:"MyGame"`
}
type Device struct {
	DeviceType string `json:"deviceType"`
	Os         string `json:"os"`
	OsVersion  string `json:"osVersion"`
}
type App1 struct {
	Version             string `json:"version"`
	DistributionChannel string `json:"distributionChannel"`
	Tag                 string `json:"tag"`
}
type Param struct {
	Leaderboard string `json:"leaderboard"`
	Amount      string `json:"amount"`
	Offset      string `json:"offset"`
}



type Apps1 struct {
	Data2   Data2   `json:"data"`
}
type Data2 struct {
	AccessToken string `json:"accessToken"`
	ClientTime  string `json:"clientTime"`
	Param       interface{}  `json:"param"`
}
func main()  {
	jsonStr := `{ "sys":{
	"appId": "angry_bird_blast_qq",
	"MyCompany": "Mybo", 
	"MyGame": "小鸟" 
    },
  "device":{
  	"deviceType": "ipad", 
  	"os": "ios", 
  	"osVersion": "1.0.0" 
  }, 
 "app": { 
 	"version": "1.1.1", 
 	"distributionChannel": "app store",
 	"tag":"{\"74\":\"1554961747\",\"80\":\"1554968154\",\"76\":\"1554989187\"}"
 },
 "data":"212213dsds"
}`
	jsonStr2 := `
{ "sys":{
	"appId": "angry_bird_blast_qq",
	"MyCompany": "Mybo", 
	"MyGame": "小鸟" 
    },
  "device":{
  	"deviceType": "ipad", 
  	"os": "ios", 
  	"osVersion": "1.0.0" 
  }, 
 "app": { 
 	"version": "1.1.1", 
 	"distributionChannel": "app store",
 	"tag":"{\"74\":\"1554961747\",\"80\":\"1554968154\",\"76\":\"1554989187\"}"
 },
 "data":{
    "accessToken":"1tgc69qwfvprlywzsdmwqlkwef6tnlhu8vpzng69zsw5q9y3869mzcveq8ftfzh9jzdns9d6kcgklx66m9t7sahr9dnsgjrh58apfb3gckysj4m3bszxayf5jvxkzcud",
    "clientTime": "1527565570",
    "param":[
     "daily_day_2018-55",
      "1",
      "0"
    ]
}
}
`
	var config Apps
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==111")
		fmt.Println(config)
		fmt.Println(config.Data2)
	}else {
		panic(err)
	}
	var config1 Apps1
	if err := json.Unmarshal([]byte(jsonStr2), &config1); err == nil {
		fmt.Println("================json str 转struct==2222")
		fmt.Println(config1)
		fmt.Println(config1.Data2.Param.([]interface{})[2])
	}else {
		panic(err)
	}
	 test := make(map[string]interface{})
	test["accessToken"] = "2dsdsds"
	test["clientTime"] = "dsadasd"
	test["param"] = "dsadasd"
	var data Data2
    var b []byte
	var err  error
	if b, err = json.Marshal(test); err == nil {
		fmt.Println("================json str 转struct==33333")
		fmt.Println(string(b))
	}else {
		panic(err)
	}

	if err := json.Unmarshal(b, &data); err == nil {
		fmt.Println("================json str 转struct==4444")
		fmt.Println(data.Param,"21212")
	}else {
		panic(err)
	}

	v4 := reflect.ValueOf(test)
	for _, k := range v4.MapKeys() {
		fmt.Println(k.String(), v4.MapIndex(k).Interface())
	}
}
