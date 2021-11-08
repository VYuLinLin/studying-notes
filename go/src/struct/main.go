package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	makeStruct()

	selectorNotation()

	changeStruct()

	tagTypeExample()

	inheritExample()

	setNameExample()

	methodsExample()

}

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func makeStruct() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "chris"

	fmt.Println(ms.i1)           // 10
	fmt.Println(ms.f1)           // 15.5
	fmt.Println(ms.str)          // chris
	fmt.Println(ms)              // &{10 15.5 chris}
	fmt.Println("ms 的内存地址", &ms) // 0xc000006028
}

type myStruct struct{ i int }

func selectorNotation() {
	var v = &myStruct{2}
	var p *myStruct
	var s myStruct

	s.i = 333
	fmt.Println("v = ", *v, "p = ", p) // v =  &{2} p =  <nil>
	fmt.Printf("myStruct = %v \n", s)  // myStruct = {333}
}

// Person name
type Person struct {
	firstName string
	lastName  string
}

func upCase(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func changeStruct() {

	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upCase(&pers1)
	fmt.Println("pers1", pers1) // pers1 {CHRIS WOODWARD}

	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward1"
	(*pers2).lastName = "Woodward2" // 这是合法的
	upCase(pers2)
	fmt.Println("pers2", pers2) // pers2 &{CHRIS WOODWARD2}

	pers3 := &Person{"Chris", "Woodward"}
	upCase(pers3)
	fmt.Println("pers3", pers3) // pers3 &{CHRIS WOODWARD}
}

// TagType s
type TagType struct {
	field1 bool   "一个布尔值"
	field2 string "介绍信息"
	field3 int    "多少钱"
}

// 带标签的结构体
func tagTypeExample() {
	tt := TagType{true, "Chris", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, i int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(i)
	fmt.Println(ixField.Index)
	fmt.Println(ixField.Name)
	fmt.Println(ixField.Type)
	fmt.Println(ixField.Tag)
}

type innerS struct {
	in1 int
	in2 int
}
type innerSs struct {
	in12 int
	in22 int
}
type outerS struct {
	b      int
	c      float32
	int    // 匿名
	innerS // 匿名
	innerSs
}

func inheritExample() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 20

	fmt.Println(outer) // &{6 7.5 60 {5 20} {0 0}}

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}, innerSs{5, 10}}
	fmt.Println(outer2) // {6 7.5 60 {5 10} {5 10}}
}

func setNameExample() {
	type A struct{ a int }
	type B struct{ a, b int }
	type C struct {
		A
		B
	}
	var c C
	fmt.Println(c.b)   // {{0} {0 0}}
	fmt.Println(c.B.a) // 0
	fmt.Println(c.A.a) // 0
}

type twoInts struct {
	a, b int
}

func (tn *twoInts) AddThem() int {
	return tn.a + tn.b
}
func (tn *twoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

type intVector []int

func (v intVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}
func methodsExample() {
	two1 := new(twoInts)
	two1.a = 12
	two1.b = 10

	fmt.Println(two1.AddThem())      // 22
	fmt.Println(two1.AddToParam(20)) // 42

	two2 := twoInts{3, 4}
	fmt.Println(two2.AddThem()) // 7

	fmt.Println(two1, two2)                                 // &{12 10} {3 4}
	fmt.Println(reflect.TypeOf(two1), reflect.TypeOf(two2)) // *main.twoInts main.twoInts

	fmt.Println(intVector{1, 2, 3}.Sum()) // 6
	fmt.Println(time.Time(time.Now()).String())
	fmt.Println(time.Now().String()[0:3])
}
