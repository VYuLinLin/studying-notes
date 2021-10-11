package main

import (
	f "fmt"

	"./sort"
)

type Namer interface {
	Method1() string
	Method2() string
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	var ai Namer

	f.Println(ai) // <nil>

	sq1 := &Square{5}
	r := Rectangle{5, 3}

	var areaIntf Shaper = sq1
	// 或者
	areaIntfs := Shaper(r)

	f.Println(sq1.Area())       // 25
	f.Println(areaIntf.Area())  // 25
	f.Println(areaIntfs.Area()) // 25

	shapers := []Shaper{r, sq1}

	for i, _ := range shapers {
		f.Println(shapers[i])
		f.Println(shapers[i].Area())
	}

	// 接口的类型断言
	if t, ok := areaIntf.(*Square); ok {
		f.Println("areaIntf的类型是Square", t) //  &{5}
	}
	if u, ok := areaIntfs.(Rectangle); ok {
		f.Println("areaIntf的类型是Rectangle", u) // {5 3}
	} else {
		f.Println("areaIntf的类型不包含Shaper")
	}

	// 使用type-switch检测
	switch t := areaIntf.(type) {
	case *Square:
		f.Println("areaIntf的类型是Square", t) // &{5}
	case *Rectangle:
		f.Println("areaIntf的类型是Rectangle", t)
	case nil:
		f.Println("areaIntf的类型是nil", t)
	default:
		f.Println("没有类型", t)
	}

	// 接口的实际应用示范
	ints()

	strings()

	days()

	anyExample()

}

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArray(data)
	sort.Sort(a)

	if !sort.IsSorted(a) {
		panic("未能成功排序")
	}
	// [-5467984 -784 0 0 42 59 74 238 905 959 7586 7586 9845]
	f.Println("此int数组已排序", a)
}

func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := sort.StringArray(data)
	sort.Sort(a)

	if !sort.IsSorted(a) {
		panic("未能成功排序")
	}
	// 8 [ friday monday saturday sunday thursday tuesday wednesday]
	f.Println("此string数组已排序", len(a), a)
}

type day struct {
	num       int
	shortName string
	longName  string
}
type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday := &day{0, "SUN", "Sunday"}
	Monday := &day{1, "MON", "Monday"}
	Tuesday := &day{2, "TUE", "Tuesday"}
	Wednesday := &day{3, "WED", "Wednesday"}
	Thursday := &day{4, "THU", "Thursday"}
	Friday := &day{5, "FRI", "Friday"}
	Saturday := &day{6, "SAT", "Saturday"}
	data := []*day{Wednesday, Thursday, Friday, Saturday, Sunday, Monday, Tuesday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("未能成功排序")
	}
	for _, d := range data {
		f.Println(d.longName)
	}
}

var i = 5
var str = "abc"

type Person struct {
	name string
	age  int
}
type Any interface{}

func anyExample() {
	var val Any
	val = 5
	f.Println("val的值为 = ", val)
	val = str
	f.Println("val的值为 = ", val)

	pers1 := new(Person)
	pers1.name = "victor"
	pers1.age = 34
	val = pers1
	f.Println("val的值为 = ", val)
}
