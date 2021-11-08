package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// 设置salary
func (p *Person) SetSalary(sal float64) {
	// p.salary = sal
	p.chF <- func() {
		p.salary = sal
	}
	if sal == 10 {
		close(p.chF)
	}
}

// 获取salary
func (p *Person) GetSalary() float64 {
	fChan := make(chan float64)
	defer close(fChan)
	p.chF <- func() { fChan <- p.salary }
	// return p.salary
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " +
		strconv.FormatFloat(p.GetSalary(), 'f', 2, 64)
}

// 使用 Channel 来并发读取对象
// 通过 channel 解决高并发修改和访问对象数据
func main() {
	bs := NewPerson("Bill", 2500.5)
	fmt.Println(bs)

	bs.SetSalary(8888.66)
	fmt.Println(bs)
	fmt.Println(bs.salary)
	for i := 1; i < 100; i++ {
		go bs.SetSalary(float64(i))
		// go bs.SetSalary(float64(i) * bs.salary)
		// fmt.Println(float64(i) * bs.salary)
		fmt.Println(bs)
	}
}
