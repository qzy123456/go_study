package main
import (
	"encoding/json"
	"os"
)
import "fmt"

type ResInfo struct {  //定义一个struct，然后这个struct里面有哪些子对象
	Data YearDataStruct
	Msg  string
}

type YearDataStruct struct {
	MouthAll []MouthStruct  //定义一个类型为数组的对象，然后这个数组的元素类型为某种struct
	Sum      DetailStruct
	Average  DetailStruct
	Quarter  []QuarterStruct
}
type DetailStruct struct {
	One   int
	Two   int
	Three int
}
type QuarterStruct struct {
	DetailStruct   //可以嵌套复合其他类型的struct，这样就继承下了其他struct的子对象
	QuarterNum int
}

type MouthStruct struct {
	Mouth        int
	PartmentItem []ItemArrStruct
}
type ItemArrStruct struct {
	PartMent string
	DetailStruct
}

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}
type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}
type Object []interface{}

func main() {
	var jsonBlob = []byte(`{"Results":[
	{"Name": "Platypus", "Order": "Monotremata","id":111},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]}`)
	//struct嵌套struct
	type Animal struct {
		Results []struct {
			Name  string
			Order string
		}
	}
	var animals Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
	jsonStr := `{"host": "http://localhost:9090",
"port": 9090,"analytics_file": "",
"static_file_version": 1,"static_dir": 
"E:/Project/goTest/src/","templates_dir": 
"E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340",
"serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	jsonStr = `{"accessToken":"507b5e08ee444dck887b66bd08672905",
"clientToken":"64e3a5415bfe405d9485f1jf2ea5c68e",
"selectedProfile":{"id":"selID","name":"Bluek404"},
"availableProfiles":[{"id":"测试ava","name":"Bluek404"}]}`
	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)

		mapTmp := dat["selectedProfile"].(map[string]interface {})
		fmt.Println(mapTmp["id"])
		/*
		var dat2 map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &dat2); err == nil {
			fmt.Println( dat2["firstName"])
		}
		*/

		mapTmp2 := (dat["availableProfiles"].([]interface {}))[0].(map[string]interface {})
		//mapTmp3 := mapTmp2[0].(map[string]interface {})
		fmt.Println(mapTmp2["id"])
	}

	//json str 转struct
	var config ConfigStruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(config)
		fmt.Println(config.Host)
	}

	//json str 转struct(部份字段)
	var part Other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(part)
		fmt.Println(part.SerTcpSocketPort)
	}

	//struct 到json str
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))
	}

	//map 到json str
	fmt.Println("================map 到json str=====================")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)

	//array 到 json str
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}

	//json 到 []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(wo)
	}

}
