package main

import (
    "fmt"
    "reflect"
)

func main() {
    //var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap := make(map[string]string)

    /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "Paris"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"

    /*使用键输出地图值 */
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    /*查看元素在集合中是否存在 */
    captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(captial) */
    /*fmt.Println(ok) */
    if ok {
        fmt.Println("美国的首都是", captial)
    } else {
        fmt.Println("美国的首都不存在")
    }
    //map
    scene := make(map[string]int)
    scene["route"] = 66
    scene["brazil"] = 4
    scene["china"] = 960
    delete(scene, "brazil")
    for k, v := range scene {
        fmt.Println(k, "=>",v)
    }

    //为了节省内存，空结构 "节省" 内存，⽐比如⽤用来实现 set 数据结构，或者实现没有 "状态" 只有⽅方法的 "静 态类"
    var null struct{}
    set := make(map[string]struct{})
    set["a"] = null
    fmt.Println(set)
    //不能同时嵌⼊入某⼀一类型和其指针类型，因为它们名字相同。
    u := Use{
        Resoure{1},
        "1121",
    }
   fmt.Println(u.id,u.Resoure.id,u.name)
    var uu Resoure = u.Resoure
    fmt.Println(uu)
    //转string
    m := Manager{Users{1,"saa" }}
    fmt.Printf("Manager: %p\n", &m)
    fmt.Println(m.ToString())
    fmt.Println(reflect.TypeOf(m.id))

    var p *Data = nil
    p.TestPointer()
    (*Data)(nil).TestPointer()  // method value
    (*Data).TestPointer(nil)    // method expression
     //p.TestValue()            //报错 invalid memory address or nil pointer dereference
    // (Data)(nil).TestValue()  //报错 cannot convert nil to type Data
    // Data.TestValue(nil)      //报错 cannot use nil as type Data in function argument
}

type Resoure struct {
    id int
}
type Use struct {
    Resoure
    name string
}
type Users struct {
    id int
    name string
}
type Manager struct {
    Users
}
func (self *Users) ToString() string {
    return fmt.Sprintf("%v", self)
}
type Data struct{}
func (Data) TestValue()    {}
func (*Data) TestPointer() {}