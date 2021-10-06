package main

import (
	"html/template"
	"os"
)

func main() {
	tEmpty := template.New("template test")
	// if 是一个空管道时的内容
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} will not print. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("Template test")
	// 如果条件满足,则为非空管道
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	// 如果条件满足,则为非空管道
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if ``}} Print IF part. {{else}} Print ELSE part. {{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

	// 如果管道为空,就会跳过with到end之前的任何内容
	t := template.New("test")
	// t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")
	// t.Execute(os.Stdout, nil)

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}} {{end}}!\n")
	t.Execute(os.Stdout, nil)

	// 模板变量 (变量名称只能由字母、数字、下划线组成。)
	tVar := template.New("test")
	tVar = template.Must(tVar.Parse("{{with $3 := `go`}}{{$3}} {{.}} {{$3}} {{end}}!\n"))
	tVar.Execute(os.Stdout, nil) // go go go !

	// range-end
	tRange := template.New("test")
	tRange = template.Must(tRange.Parse("{{range .}}\n{{.}}{{end}}\n"))
	s := []int{1, 2, 3, 4}
	tRange.Execute(os.Stdout, s)

	// 预定义模板函数
	preT := template.New("test")
	preT = template.Must(preT.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	preT.Execute(os.Stdout, nil) // hello Mary!
}
