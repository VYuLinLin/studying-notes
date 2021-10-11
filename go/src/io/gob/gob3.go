package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Address struct {
	Type int64
	City string
	// Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	// Remark    string
}

func main() {
	file, _ := os.OpenFile("vcard.gob", os.O_RDONLY, 0)
	defer file.Close()
	fmt.Println(file) // &{0xc00007e780}

	var v VCard
	dec := gob.NewDecoder(file)
	dec.Decode(&v)

	fmt.Printf("%#v \n", v)                  // main.VCard{FirstName:"Jan", LastName:"Kersschot", Addresses:[]*main.Address{(*main.Address)(0xc000004828), (*main.Address)(0xc000004840)}}
	fmt.Printf("%v \n", v.Addresses[0].Type) // 1
	fmt.Printf("%v \n", v.Addresses[0].City) // Aartselaar
	fmt.Printf("%v \n", v.Addresses[1].Type) // 2
	fmt.Printf("%v \n", v.Addresses[1].City) // Boom
}
