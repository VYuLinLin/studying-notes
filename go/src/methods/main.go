package main

import (
	"fmt"
	"runtime"
)

type Log struct {
	msg string
}

// 聚合（或组合）
type Customer struct {
	Name string
	Log
}

func main() {
	c := new(Customer)
	c.Name = "Victor"
	c.Log = Log{}
	c.Log.msg = "1 - yes we can!"

	fmt.Println(c) // &{Victor {1 - yes we can!}}

	c = &Customer{"Vector", Log{"2 - yes we can!"}}

	fmt.Println(c.Log.msg) // 2 - yes we can!
	fmt.Println(c.Log)     // &{2 - yes we can!}
	fmt.Println(c.Logs())  // 2 - yes we can!

	c.Log.Add("在我之后，世界会变得更美好")
	fmt.Println(c.Log.msg) // 在我之后，世界会变得更美好

	fmt.Println(c.Log)

	extendExample()
}

func (l Log) Logs() string {
	return l.msg
}
func (l *Log) Add(s string) {
	l.msg = "\n" + s
}

// func (l *Log) String() string {
// 	return l.msg
// }

// func (c *Customer) Log() *Log {
// 	return c.log
// }

// ====================================
type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

type TT float32

func (t TT) String() string {
	return fmt.Sprintf("%v", t)
}
func extendExample() {
	v := new(Voodoo)
	v.Magic() // voodoo magic
	v.MoreMagic()

	// t := new(TT)
	// t.String()

	// 查看当前内存情况
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024) // 127 Kb
}
