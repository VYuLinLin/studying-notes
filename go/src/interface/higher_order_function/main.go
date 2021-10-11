package main

import f "fmt"

type Any interface{}
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}
type Cars []*Car

func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	i := 0
	cs.Process(func(c *Car) {
		result[i] = f(c)
		i++
	})
	return result
}
func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			var cars = sortedCars[c.Manufacturer]
			sortedCars[c.Manufacturer] = append(cars, c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
func main() {
	// make some cars:
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	// query
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	f.Printf("全部汽车：%v\n", allCars)
	f.Printf("新的宝马车：%#v\n", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	BMWCount := len(sortedCars["BMW"])

	f.Println("排序后的汽车集合：", sortedCars)
	f.Println("我们现在有", BMWCount, "辆宝马车")
	for _, v := range sortedCars["BMW"] {
		f.Println("拥有的宝马车是", v)
	}
}
