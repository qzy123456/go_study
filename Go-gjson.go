package main
import (
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/gjson"
)

const jsons = `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":["Sara","Alex","Jack"],"fav.movie":"Deer Hunter","friends":[{"first":"Dale","last":"Murphy","age":44},{"first":"Roger","last":"Craig","age":68},{"first":"Jane","last":"Murphy","age":47}]}`

func main() {
	// 首先我们判断该json是否合法
	if !gjson.Valid(jsons) {
		log.Fatalf("%s", "invalid json")
	}
	// 获取Json中的age
	age := gjson.Get(jsons, `age`).Int()
	fmt.Printf("%T, %+v\n", age, age)
	// 获取lastname
	lastname := gjson.Get(jsons, `name.last`).String()
	fmt.Printf("%T, %+v\n", lastname, lastname)
	// 获取children数组
	for _, v := range gjson.Get(jsons, `children`).Array() {
		fmt.Printf("%q ", v.String())
	}
	fmt.Println()
	// 获取第二个孩子
	fmt.Printf("%q\n", gjson.Get(jsons, `children.1`).String())
	fmt.Printf("%q\n", gjson.Get(jsons, `children|1`).String())
	// 通配符获取第三个孩子
	fmt.Printf("%q\n", gjson.Get(jsons, `child*.2`).String())
	// 反转数组函数
	fmt.Printf("%q\n", gjson.Get(jsons, `children|@reverse`).Array())
	// 自定义函数 - 全转大写
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}
		return json
	})
	fmt.Printf("%+v\n", gjson.Get(jsons, `children|@case:upper`).Array())
	// 直接解析为map
	jsonMap := gjson.Parse(jsons).Map()
	fmt.Printf("%+v\n", jsonMap)
	for _, v := range jsonMap {
		fmt.Printf("%T, %+v\n", v, v)
	}
}
