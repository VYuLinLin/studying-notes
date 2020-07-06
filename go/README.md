# GO语言学习笔记
学习时的Go语言版本1.14.4
Go语言起源于2007年，由Google公司研发与维护，于2009年正式对外发布。
学习地址：

[Go指南](http://go-tour-zh.appspot.com/list)

[Go 入门指南|技术论坛](https://learnku.com/docs/the-way-to-go)

Go语言的特性：

- Go语言没有类和继承的概念，它通过接口（interface）的概念来实现多态性。
- Go语言属于强类型、静态类型语言。
- 函数是Go语言中的基本构件。
- Go语言支持交叉编译。比如在Linux系统上开发Windows下运行的应用程序。

基础指令：

- go version    查看Go版本信息
- gofmt hello.world.go      打印格式化后的结果而不重写文件
- gofmt -w hello.world.go   格式化并重写原代码
- gofmt -w *.go     会格式化并重写所有 Go 源文件
- gofmt map1        会格式化并重写 map1 目录及其子目录下的所有 Go 源文件
- gofmt -r '(a) -> a' –w *.go       替换所有源文件中的(a) 为 a

目前版本实现的三个Go工具

1. go install       是安装Go包的工具，主要安装非标准库的包文件，将源文件编译成对象文件。
2. go fix       将Go代码从旧的发行版迁移到最新的发行版
3. go test      是一个轻量级的单元测试框架

---

Go语言注意事项：

1. 不能使用单个下划线（_）作为标识符。
2. 字符串不能使用单引号。



