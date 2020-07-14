# GO语言学习笔记
学习时的Go语言版本1.14.4
Go语言起源于2007年，由Google公司研发与维护，于2009年正式对外发布。
学习地址：

[Go指南](http://go-tour-zh.appspot.com/list)

[Go 入门指南|技术论坛](https://learnku.com/docs/the-way-to-go)

Go的格言： “没有不必要的代码！”

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

Go语言数据的基本类型：
1. bool     布尔类型
2. string   字符串类型
3. number   数字类型
    1. int      整数    是计算最快的类型
        1. int8     -128 -> 127
        2. int16    -32768 -> 32768
        3. int32    -2,147,483,648 -> 2,147,483,647
        4. int64    -9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807
    2. uint     无符号整数
        1. uint8    0 -> 255
        2. uint16   0 -> 65535
        3. uint32   0 -> 4,294,967,295
        4. uint64   0 -> 18,446,744,073,709,551,615
    3. float    浮点型
        1. float32  +- 1e-45 -> +- 3.4 * 1e38   精确到小数点后7位
        2. float64  +- 5 1e-324 -> 107 1e308    精确到小数点后15位，推荐优先使用

---

Go语言注意事项：

1. 不能使用单个下划线（_）作为标识符，实际上下划线（_）是一个只写变量。
2. 字符串不能使用单引号。
3. 每个 Go 文件都属于且仅属于一个包
4. 每一段代码只会被编译一次
5. Go语言中包中的对象，以是否大小写开头来区分是否被导出的对象，其中大写开头的标识符对象为导出（类似于面向对象语言中的public），其中小写开头的标识符的对象对外是不可见的，对内为可见且可用的（类似于面向对象语言中的private）。
6. Go语言的作用域中，只有函数提升，没有变量提升
7. 自定义函数无法用于常量的复制，因为在编译期间自定义函数属于未知，但是内置函数可以，如 len()。
8. 变量的初始化声明的简写方式（:=），系统会自动推断类型，只能在函数体内使用。
9. 局部作用域声明的变量必须使用，全局声明的变量可以不使用。
10. 多变量赋值只能在局部作用域使用，多变量声明不受此限制。
11. 条件判断时，哪怕只有一行代码，也不能省略大括号{}
12. for 和 if 等条件判断不能用()

---

基本知识：

- 使用import导入包后，可以定义或声明0或多个常量（const）、变量（var）、类型（type）
- 使用关键字 func 定义函数
- main 函数是每一个可执行程序所必须包含的，每个包中只能有一个main函数。
- init和main函数都是自执行函数，默认情况下执行顺序依次是init > main，并且都不能有参数和返回值。
- go doc    命令可以显示包的文档说明
- 常量的数据类型只能是：布尔型、数字型（整数型、浮点型和复数）、字符串型
- 常量支持并行赋值（类似于JavaScript中的解构赋值）
- 关于iota，简单理解就是，每遇到一次const关键字，iota就重置为0。
- 当一个变量被声明后，系统会自动赋给该变量类型的零值：
    - int       0
    - float     0.0
    - bool      false
    - string    ""
    - 指针      nil
- 所有的内存在Go中都是经过初始化的。
- rand.Intn(n)     返回介于[0, n]之间的伪随机数，不包括n
- rand.Float32()和rand.Float64()返回介于[0.0, 1.0]之间的数字，包括0.0，不包括1.0
- rand.Seed(value)  提供伪随机数的生成种子，推荐使用当前时间的纳秒数
- strings   以下是strings包中关于字符串操作的函数
    1. HasPrefix(s, prefix string) bool 判断字符串 s 是否以 prefix 开头
    2. HasSuffix(s, suffix string) bool 判断字符串 s 是否以 suffix 结尾
    3. Contains(s, substr string) bool  判断字符串 s 是否包含 substr
    4. Index(s, str string) int 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str
    5. LastIndex(s, str string) int 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str
    6. IndexRune(s string, r rune) int  如果 ch 是非 ASCII 编码的字符，建议使用此函数来对字符进行定位
    7. Replace(str, old, new, n) string 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new
    8. Count(s, str string) int 计算字符串 str 在字符串 s 中出现的非重叠次数
    9. Repeat(s, count int) string  用于重复 count 次字符串 s 并返回一个新的字符串
    10. ToLower(s) string  将字符串中的 Unicode 字符全部转换为相应的小写字符
    11. ToUpper(s) string  将字符串中的 Unicode 字符全部转换为相应的小写字符
    12. TrimSpace(s) string 剔除字符串开头和结尾的空白符号.
    13. Trim(s, "cut") string   将开头和结尾的 cut 去除掉
    14. TrimLeft(s, 'cut') string   只将开头的 cut 去除掉
    14. TrimRight(s, 'cut') string   只将结尾的 cut 去除掉
    15. Fields(s) slice 利用空白作为分隔符，分割成若干份
    16. Split(s, sep)  slice    自定义分割符号对字符串分割
    17. strings.Join(sl []string, sep string) string    将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
    18. NewReader(str) 生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针
- strconv   以下是strconv包中关于字符串与其他类型转换的函数
    1. IntSize    获取程序运行的操作系统平台下 int 类型所占的位数
    2. Itoa(i int) string   返回数字 i 所表示的字符串类型的十进制数
    3. FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
    4. Atoi(s string) (i int, err error)    将字符串转换为 int 型。
    5. ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。

- time 提供了一个数据类型 time.Time（作为值使用）以及显示和测量时间和日期的功能函数
    1. Now()    获取当前时间，如 var t = time.Now()
    2. t.Day()      天
    3. t.Month()    月
    4. t.Year()     年
    5. t.Format(layout string) string   根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串，预定义格式time.ANSIC 或 time.RFC822

- 一个指针变量可以指向任何一个值的内存地址
- 指针未分配到任何变量时，值为 nil ，一般指针变量缩写为 ptr
- &a（返回变量a的内存地址）
- *int（声明指针变量的类型）
- *a（返回变量a的值）

————————————————
优先级     运算符

 7      ^ !

 6      * / % << >> & &^

 5      + - | ^

 4      == != < <= >= >

 3      <-

 2      &&

 1      ||

————————————————
