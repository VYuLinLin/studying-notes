package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "beijing", "china"}
	wa := &Address{"work", "shenzhen", "china"}
	vc := VCard{"Chris", "Wang", []*Address{pa, wa}, "none"}

	js, _ := json.Marshal(vc)
	var sourceData VCard
	json.Unmarshal(js, &sourceData)

	// 以下两种打印，效果相同，值不同
	fmt.Println(vc)         // {Chris Wang [0xc000076480 0xc0000764b0] none}
	fmt.Println(sourceData) // {Chris Wang [0xc000076810 0xc000076840] none}
	// 以下两种打印，效果相同，值不同
	fmt.Printf("JSON format: %s\n", js)
	fmt.Println("\n格式化后的json", string(js))

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	// or
	// err = enc.Encode(sourceData)
	if err != nil {
		log.Println("Error in encoding json")
	}

	// 解码任意的数据
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	var f interface{}
	err = json.Unmarshal(b, &f)

	fmt.Printf("\n解码任意数据%s", f)
	fmt.Println("\n解码任意数据", f)
}
