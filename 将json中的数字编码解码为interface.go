package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// 将 decode 的值转为 int 使用
func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	var status = uint64(result["status"].(float64))
	fmt.Println("Status value: ", status)
	
	//////////////////////////////////////
	test1()
	//////
	test2()
	//////
	test3()
	/////
	test4()
	////// slice 切割
	test5()
	
}
//使用 Decoder 类型来 decode JSON 数据，明确表示字段的值类型
func test1()  {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}

	var status, _ = result["status"].(json.Number).Int64()
	fmt.Println("Status value: ", status)
}
//// // 你可以使用 string 来存储数值数据，在 decode 时再决定按 int 还是 float 使用
// // 将数据转为 decode 为 string
func test2()  {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}
	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&result); err != nil {
		log.Fatalln(err)
	}
	var status uint64
	err := json.Unmarshal([]byte(result["status"].(json.Number).String()), &status);
	println(err)
	fmt.Println("Status value: ", status)
}
//使用 struct 类型将你需要的数据映射为数值型
func test3()  {
	var data = []byte(`{"status": 200}`)
	var result struct {
		Status uint64 `json:"status"`
	}

	err := json.NewDecoder(bytes.NewReader(data)).Decode(&result)
	_ =err
	fmt.Printf("Result: %+v", result,"\n")
}
//可以使用 struct 将数值类型映射为 json.RawMessage 原生数据类型
//适用于如果 JSON 数据不着急 decode 或 JSON 某个字段的值类型不固定等情况：
func test4()  {
	records := [][]byte{
		[]byte(`{"status":200, "tag":"one"}`),
		[]byte(`{"status":"ok", "tag":"two"}`),
	}

	for idx, record := range records {
		var result struct {
			StatusCode uint64
			StatusName string
			Status     json.RawMessage `json:"status"`
			Tag        string          `json:"tag"`
		}

		err := json.NewDecoder(bytes.NewReader(record)).Decode(&result)
		_=err

		var name string
		err = json.Unmarshal(result.Status, &name)
		if err == nil {
			result.StatusName = name
		}

		var code uint64
		err = json.Unmarshal(result.Status, &code)
		if err == nil {
			result.StatusCode = code
		}

		fmt.Printf("[%v] result => %+v\n", idx, result)
	}
}

func test5()  {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // 4
	dir1 := path[:sepIndex:sepIndex]		// 此时 cap(dir1) 指定为4， 而不是先前的 16
	dir2 := path[sepIndex+1:]
	dir1 = append(dir1, "suffix"...)

	path = bytes.Join([][]byte{dir1, dir2}, []byte{'/'})
	println("dir1: ", string(dir1))		// AAAAsuffix
	println("dir2: ", string(dir2))		// BBBBBBBBB
	println("new path: ", string(path))	// AAAAsuffix/BBBBBBBB
}