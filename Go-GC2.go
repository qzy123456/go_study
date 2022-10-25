package main

import (
	"log"
	"runtime"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Close() {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Close")
}

func (p *Person) NewOpen() {
	log.Println("Init")
	runtime.SetFinalizer(p, (*Person).Close)
}

func Tt(p *Person) {
	p.Name = "NewName"
	log.Println(p)
	log.Println("Tt")
}

// 查看内存情况
func Mem(m *runtime.MemStats) {
	runtime.ReadMemStats(m)
	log.Printf("%d Kb\n", m.Alloc/1024)
}

func main() {
	var m runtime.MemStats
	Mem(&m)

	var p *Person = &Person{Name: "lee", Age: 4}
	p.NewOpen()
	log.Println("Gc完成第一次")
	log.Println("p:", p)
	runtime.GC()
	time.Sleep(time.Second * 5)
	Mem(&m)

	var p1 *Person = &Person{Name: "Goo", Age: 9}
	runtime.SetFinalizer(p1, Tt)
	log.Println("Gc完成第二次")
	time.Sleep(time.Second * 2)
	runtime.GC()
	time.Sleep(time.Second * 2)
	Mem(&m)

}