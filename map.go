package main

import (
    "fmt"
    "reflect"
    "sort"
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
    duoWeiMap()
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

func duoWeiMap(){
        m := make(map[int]string) //初始化一个map
        m[1] = "ok"//赋值
        delete(m,1)//删除m键为1的键值对
        a := m
        fmt.Println(a)
        //m2 := make(map[int]map[int]string) //初始化一个二维map
        //m2[1][1] = "ok2"
        //fmt.Println(m2)//在这里m2运行不了，报错信息为panic: assignment to entry in nil map
        var m2 map[int]map[int]string//定义一个二维数组的类型
        m2 = make(map[int]map[int]string) //初始化这个map,定义类型与初始化概念
        m2[1] = make(map[int]string)//初始化内部数组，如果不初始化会报错，也就是说多维数组需要把内部map全部初始化，
        //这里是将m2键为1的值初始化，之后如果有2号键，则需要对2也进行初始化
        //这就肯定会有一个解决方案，在实际编码中我们是不会手动初始化一个未知个数的map的
        m2[1][1] = "ok2"
        fmt.Println(m2)
        fmt.Println(m2[2][1]) //这里可以看出返回的是一个空字符串
        fmt.Println("结束")
        v, b := m2[2][1] //多返回值的写法，v为键对应的值，b为这个值是否存在。
        //因为有些时候我们会给值认为赋值，而这个值恰巧就是空字符串。这里返回的false意味着并没有值
        fmt.Println(v, b)
        if !b {
            m2[2] = make(map[int]string)
        }
        m2[2][1] = "" //如果不加以赋值，那么b还是为空,如果给予赋值，则为true，哪怕是人工赋值一个空字符串
        v, b = m2[2][1]
        fmt.Println(v, b)
        fmt.Println("开始range")
        for k, v := range m2 { //for配合range可以实现php类似foreach的功能。
            //但是听讲课老师所说，好像range很神奇的样子...在php里foreach岂不是更强大
            fmt.Println(k, v) //k为索引或者键，v为对应的值，循环出来的也只是一个对数据的拷贝。
            //在循环内修改k或者v的话是并不影响m2的，php的foreach同理
            //如果循环内只需要k或者v可以在用下划线代替例如_,v:=range....，这里就省略了k只保留v。
            //如果只写一个返回值，那默认的就是返回索引或者说键了，这点和foreach相反
        }
        m3 := make([]map[int]string, 5) //使用make创建一个slice,长度为5，golang推荐用make创建。
        //再回顾一下，第二个参数为长度，第三个参数为容量(可省略)
        //m3代表创建了一个类型为map的slice,slice键为索引，值为map类型
        fmt.Println(m3)
        for k := range m3 {
            m3[k] = make(map[int]string) //将slice中每一个值都初始化,
            m3[k][1] = "直直"//将这个map索引为1的赋值
        }
        fmt.Println(m3)
        fmt.Println("如何对map进行排序")//map的无序性
        m4 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
        for k, v := range m4 {
            fmt.Println(k, v)
        } //在这里多重复几次，可以看出并不是按照顺序来打印的
        //问题，map具有无序性，那么如果对map进行排序,思路，for循环维护一计数机和一个切片
        s := make([]int, len(m4))
        i := 0
        for k := range m4 {
            s[i] = k
            i += 1
        }
        sort.Ints(s) //引入sort包，对切片值进行排序
        fmt.Println(s)}