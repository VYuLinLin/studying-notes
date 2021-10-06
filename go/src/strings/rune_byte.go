package main

import f "fmt"

// rune_byte
func main() {
	first := "first"
	f.Println([]rune(first)) // [102 105 114 115 116]
	f.Println([]byte(first)) // [102 105 114 115 116]
	f.Println(first[:2])     // fi

	text := "天下我最帅"
	rune := []rune(text)
	byte := []byte(text)
	f.Println(rune)             // [22825 19979 25105 26368 24069]
	f.Println(byte)             // [229 164 169 228 184 139 230 136 145 230 156 128 229 184 133] 三个字节
	f.Println(string(text[:2])) // ��
	f.Println(string(text[:6])) // 天下
	f.Println(string(rune[:2])) // 天下
	f.Println(string(byte[:2])) // ��
	f.Println(string(byte[:6])) // 天下

	var a, _ = f.Printf("%s", text) // 返回所有的bytes数量
	f.Println(a)                    // 15

}
