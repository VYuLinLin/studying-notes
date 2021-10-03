# GO 语言学习笔记

学习时的 Go 语言版本 1.14.4
Go 语言起源于 2007 年，由 Google 公司研发与维护，于 2009 年正式对外发布。
学习地址：

[Go 指南](http://go-tour-zh.appspot.com/list)

[Go 入门指南|技术论坛](https://learnku.com/docs/the-way-to-go)

Go 的格言： “没有不必要的代码！”

Go 语言的特性：

- Go 语言没有类和继承的概念，它通过接口（interface）的概念来实现多态性。
- Go 语言属于强类型、静态类型语言。
- 函数是 Go 语言中的基本构件。
- Go 语言支持交叉编译。比如在 Linux 系统上开发 Windows 下运行的应用程序。

基础指令：

- ```go env``` 查看go所有的环境变量
- ```go env GOPATH``` 查看 GOPATH 的系统路径
- ```go version``` 查看 Go 版本信息
- ```go clean```  删除编译生成的可执行文件
- ```go run main.go```    先构建main.go文件里包含的程序，然后执行构建后的程序
- ```gofmt hello.world.go``` 打印格式化后的结果而不重写文件
- ```gofmt -w hello.world.go``` 格式化并重写原代码
- ```gofmt -w \*.go``` 会格式化并重写所有 Go 源文件
- ```gofmt map1``` 会格式化并重写 map1 目录及其子目录下的所有 Go 源文件
- ```gofmt -r '(a) -> a' –w \*.go``` 替换所有源文件中的(a) 为 a

依赖包管理
  1. ```go mod init project_name``` 在项目的根目录初始化包管理，自动分析依赖，创建go.mod和go.sum
  2. ```go mod tidy``` 自动分析依赖，自动添加和删除依赖
  3. ```go mod vendor``` 创建vendor目录，将依赖拷贝到当前的vendor文件中
  4. ```go mod download``` 手动下载依赖

打包构建
- ```go build main.go```在main.go所在的目录下，生成windows可执行的main.exe文件

自动化构建工具 [goreleaser](https://github.com/goreleaser/goreleaser)
  1. [下载](https://github.com/goreleaser/goreleaser/releases)
  2. 在项目根目录使用```goreleaser init```初始化配置文件
  3. 根据需要修改 .goreleaser.yml 文件，或保持默认配置
  4. ```goreleaser --snapshot --skip-publish --rm-dist``` 打包并跳过发布
  
目前版本实现的三个 Go 工具

1. go install 是安装 Go 包的工具，主要安装非标准库的包文件，将源文件编译成对象文件。
2. go fix 将 Go 代码从旧的发行版迁移到最新的发行版
3. go test 是一个轻量级的单元测试框架

---

Go 语言数据类型：
- 基本类型
1. bool 布尔类型
2. string 字符串类型
3. number 数字类型
   1. int 整数 是计算最快的类型
      1. int8 -128 -> 127
      2. int16 -32768 -> 32768
      3. int32 -2,147,483,648 -> 2,147,483,647
      4. int64 -9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807
   2. uint 无符号整数
      1. uint8 0 -> 255
      2. uint16 0 -> 65535
      3. uint32 0 -> 4,294,967,295
      4. uint64 0 -> 18,446,744,073,709,551,615
   3. float 浮点型
      1. float32 +- 1e-45 -> +- 3.4 \* 1e38 精确到小数点后 7 位
      2. float64 +- 5 1e-324 -> 107 1e308 精确到小数点后 15 位，推荐优先使用


- 结构化（复合）类型
1.  struct（ 结构体 ） 值类型，不限类型的属性和值组成
2. array （ 数组 ） 值类型，长度固定，类型统一，初始化 [3]int
3. slice （ 切片） 引用类型，长度可变，类型统一，对数组的一个连续片段的引用。初始化 []int
4. Map 引用类型，无序，一个由 key 和 value 组成的类型统一的键值对结构数据
5. channel （ 管道 ）

- 描述类型
1. interface （ 接口 ）

根据存储类型又分为值类型和引用类型：
- 值类型：int系列、float系列、bool、string、数组和结构体

  特点：变量直接存储值，内存通常在栈中分配
- 引用类型：指针、slice切片、管道channel、接口interface、map、函数等

  特点：变量存储的是一个地址，这个地址对应的空间里才是真正存储的值，内存通常在堆中分配
---

Go 语言注意事项：

1. 不能使用单个下划线（ \_ ）作为标识符，实际上下划线（ \_ ）是一个只写变量。
2. 字符串不能使用单引号。
3. 每个 Go 文件都属于且仅属于一个包
4. 每一段代码只会被编译一次
5. Go 语言中包中的对象，以是否大小写开头来区分是否被导出的对象，其中大写开头的标识符对象为导出（类似于面向对象语言中的 public），其中小写开头的标识符的对象对外是不可见的，对内为可见且可用的（类似于面向对象语言中的 private）。
6. Go 语言的作用域中，只有函数提升，没有变量提升
7. 自定义函数无法用于常量的复制，因为在编译期间自定义函数属于未知，但是内置函数可以，如 len()。
8. 变量的初始化声明的简写方式（:=），系统会自动推断类型，只能在函数体内使用。
9. 局部作用域声明的变量必须使用，全局声明的变量可以不使用。
10. 多变量赋值只能在局部作用域使用，多变量声明不受此限制。
11. 条件判断时，哪怕只有一行代码，也不能省略大括号{}
12. for 和 if 等条件判断不能用()
13. 函数定义时，只要有命名返回值，必须使用()括起来
14. new() 是一个函数，不要忘记它的括号
15. 匿名函数必须在函数作用域中声明，不能在全局作用域声明
16. 数组的数目，在声明时就需要给出（编译时需要知道数组的长度以便分配内存），数组的长度最大为 2GB
17. 打印时强烈建议引用 fmt 包，内置函数 print 和 println 打印复杂类型时要么报错，要么只打印内存地址
18. 切片不能被重新分片以获取数组的前一个元素，但是可以获取后一个元素，如果后一个元素有值
19. 绝对不要用指针指向 slice，因为切片本身是一个引用类型，所以它本身就是一个指针
20. len()函数计算字符串的长度时是按字节数计算的，但是可以通过 len([]int32(s)) 来获得字符串的数量，不过使用 utf8.RuneCountInString(s)计算的效率会更高。
21. str[start:end] 获取某个字符串， 是按字节数索引的
22. Go 语言中字符串是不可变的，如果需要更改字符串，需要先转换成字节数组，改变后在用内置函数 string()转换成字符串
23. 相等运算符（==）并不是严格意义上的相等，也不是非严格意义上的相等，如 0 == 0.00 结果是 true
24. new() 用于值类型， make() 用于引用类型
25. 在一个结构体中，每种数据类型，只能有一个匿名字段
26. 结构体中，字段名，外层名字会覆盖内层名字，但是在同一级别有两个相同的名字时，调用时会报错，必须程序员自己修正
27. 不要再 String()方法里面调用涉及 String()方法的方法，会导致意料之外的错误，比如隐式的循环调用，会导致内存溢出。
28. switch case 语法中，可以不用像 js 中那样使用 break
29. 不要通过共享内存来通信，而通过通信来共享内存
30. 不要使用全局变量或者共享内存，因为会在并发运算的时候带来不确定性
31. 通过打印状态来确定通道的发送和接口顺序是不准确的，因为打印状态和通道实际发生读写的时间延迟会导致和真实发生的顺序不同。
32. Go 语言中，没有引用传递，函数都是按照值传递，即 passed by value。如果函数调用时，参数是一个指针，则 Go 会把指针进行隐式转换得到 value，但返回来则不行。
33. Windows 系统换行用"\r\n", Linux 平台下使用 "\n"
34. defer 仅在函数返回时才会执行,在循环的结尾或其他一些有限范围的代码内不会执行
35. 当切片作为参数传递时,切记不要解引用切片。
36. 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
37. 只有当代码中并发执行非常重要，才使用协程和通道。
38. ```import "test/hello"```   这里的hello是指路径不是包名
39. import 是按文件夹名来查找，一个文件夹下的所有文件都属于同一个包，所以函数变量不能重复。
40. 文件夹命名不能包含特殊字符，如_、-、+等
41. 包名统一使用小写，不建议使用驼峰和下划线（golint）。
42. 由于 Go 语言不存在隐式类型转换，因此所有的转换都必须显式说明。
43. 在Go语言中，指针无法获取文字或常量的地址。
44. fallthrough 必须在case语句中的最后一行。
45. 关键字 continue 只能被用于for循环中
46. 函数不能在其他函数内声明，除非是匿名函数
---

基本知识：

- 使用 import 导入包后，可以定义或声明 0 或多个常量（const）、变量（var）、类型（type）
- 使用关键字 func 定义函数
- main 函数是每一个可执行程序所必须包含的，每个包中只能有一个 main 函数。
- init 和 main 函数都是自执行函数，默认情况下执行顺序依次是 init > main，并且都不能有参数和返回值。
- go doc 命令可以显示包的文档说明
- 常量的数据类型只能是：布尔型、数字型（整数型、浮点型和复数）、字符串型
- 常量支持并行赋值（类似于 JavaScript 中的解构赋值）
- 关于 iota，简单理解就是，每遇到一次 const 关键字，iota 就重置为 0。
- 当一个变量被声明后，系统会自动赋给该变量类型的零值：
  - int 0
  - float 0.0
  - bool false
  - string ""
  - 指针 nil
- 所有的内存在 Go 中都是经过初始化的。
- rand.Intn(n) 返回介于[0, n)之间的伪随机数，不包括 n
- rand.Float32()和 rand.Float64()返回介于[0.0, 1.0)之间的伪随机数，包括 0.0，不包括 1.0
- rand.Seed(value) 提供伪随机数的生成种子，推荐使用当前时间的纳秒数
- 字符串本质上是一个字节数组
- strings 以下是 strings 包中关于字符串操作的函数
  1. HasPrefix(s, prefix string) bool 判断字符串 s 是否以 prefix 开头
  2. HasSuffix(s, suffix string) bool 判断字符串 s 是否以 suffix 结尾
  3. Contains(s, substr string) bool 判断字符串 s 是否包含 substr
  4. Index(s, str string) int 返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str
  5. LastIndex(s, str string) int 返回字符串 str 在字符串 s 中最后出现位置的索引（str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str
  6. IndexRune(s string, r rune) int 如果 ch 是非 ASCII 编码的字符，建议使用此函数来对字符进行定位
  7. Replace(str, old, new, n) string 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串，如果 n = -1 则替换所有字符串 old 为字符串 new
  8. Count(s, str string) int 计算字符串 str 在字符串 s 中出现的非重叠次数
  9. Repeat(s, count int) string 用于重复 count 次字符串 s 并返回一个新的字符串
  10. ToLower(s) string 将字符串中的 Unicode 字符全部转换为相应的小写字符
  11. ToUpper(s) string 将字符串中的 Unicode 字符全部转换为相应的小写字符
  12. TrimSpace(s) string 剔除字符串开头和结尾的空白符号.
  13. Trim(s, "cut") string 将开头和结尾的 cut 去除掉
  14. TrimLeft(s, 'cut') string 只将开头的 cut 去除掉
  15. TrimRight(s, 'cut') string 只将结尾的 cut 去除掉
  16. Fields(s) slice 利用空白作为分隔符，分割成若干份
  17. Split(s, sep) slice 自定义分割符号对字符串分割
  18. strings.Join(sl []string, sep string) string 将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
  19. NewReader(str) 生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针
- strconv 以下是 strconv 包中关于字符串与其他类型转换的函数

  1. IntSize 获取程序运行的操作系统平台下 int 类型所占的位数
  2. Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数
  3. FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64。
  4. Atoi(s string) (i int, err error) 将字符串转换为 int 型。
  5. ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型。

- time 提供了一个数据类型 time.Time（作为值使用）以及显示和测量时间和日期的功能函数

  1. Now() 获取当前时间，如 var t = time.Now()
  2. t.Day() 天
  3. t.Month() 月
  4. t.Year() 年
  5. t.Format(layout string) string 根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串，预定义格式 time.ANSIC 或 time.RFC822
  6. Ticker(duration) 以指定的时间间隔重复的向通道 C 发送时间值

- 一个指针变量可以指向任何一个值的内存地址
- 指针未分配到任何变量时，值为 nil ，一般指针变量缩写为 ptr
- &a（返回变量 a 的内存地址）
- \*int（声明指针变量的类型）
- \*a（返回变量 a 在内存地址中的值）
- fallthrough 会继续switch case的判断
- break 只能结束一层循环，LABEL: break可以结束多层循环，但是label必须贴首行，加:冒号
- LABEL: break 与 LABEL: continue 用法一致
- 函数定义时，如果定义的是命名返回值，则 return 可以不带参数
- 函数即便定义了命名返回值，依旧可以无视它而返回明确的值
- params ...interface{} 可以获取到未定义类型/未定义数量的入参集合
- defer 会推迟到函数返回之前（或执行return语句之后）才执行
- 当有多个 defer 注册时，执行顺序为逆序执行（类似于栈，后进先出）
- 内置函数：

  1. close(ch) 关闭一个通道，该通道必须是双向或者仅发送的。
  2. len 返回某个类型的长度和数量（字符串、数组、切片、map、管道）
  3. cap 返回某个类型的最大容量（只能用于切片、map）
  4. new 分配内存，用于值类型和用户自定义的类型,如自定义结构。(数组、结构、值类型)
  5. make 分配内存，用于内置引用类型（切片slice、字典(map)、管道channel）
  6. copy 拷贝切片，后面的值会拷贝给前面的值，返回已拷贝的数量，如：copy(slTo, slFrom)
  7. append 从切片后面添加新元素，返回新切片 如：append(sl3, 4, 5, 6) 或者 append(sl3, sl4...)
  8. panic、recover 都用于错误处理
  9. print、println 底层打印函数，部署环境中建议用 fmt 包
  10. complex、real、imag 用于创建和操作复数

- 如果需要大量的重复计算时，建议使用数组或切片或 map 缓存计算的值，对优化计算性能，效果显著。
- 数组是具有相同唯一类型的一组已编号且长度固定的数据项序列。类型可以是原始类型或自定义类型，长度必须是一个常量表达式，且必须是一个非负整数。
- Go 语言中，数组是一种值类型，所以可以通过 new()来创建
- 想在函数中修改复杂类型的原值，需要传变量的引用指针
- 切片（slice）是对数组的一个连续片段的引用，这个引用可以是整个数组或一部分
- 切片是引用类型，且是一个长度可变的数组，最小为 0，最长为相关数组的长度
- 切片的初始化格式： var slice1 []type = arr1[start: end]
- slice2 := slice1[:cap(slice1)] 可以把切片扩展到最大
- 修改切片中的值，会影响到原数组的值
- 因为字符串是纯粹不可变的字节数组，所以他们也可以被切分成切片
- new() 和 make() 的区别
  - new(T) 返回一个指向类型为 T，值为 0 的地址的指针，适用于值类型，如数组和结构体
  - make(T) 返回一个类型为 T 的初始值，适用于引用类型，如slice、map、channel
- 类型检测
- xt := reflect.TypeOf(var) 返回检测对象的类型
- xv := reflect.ValueOf(var) 返回检测对象的值
- xt.Kind() 和 xv.Kind() 总是返回底层类型
- 其中 xv.Type() 返回类型，xt 没有 Type 方法
- xv.Interface() 返回还原后的值
- v.(type) 此种方式只能在 switch 中使用，且需提前知道有几种类型，函数没有返回值
- map 是引用类型，可以动态添加，未初始化的 map 值为 nil，声明格式：var map1 map[keytype]valuetype
- map 数据类型中的 key 一般是基本类型，value 可以是任何类型
- \_, ok := map1[key1] 可以通过 ok 的 bool 值判断 map1 中是否有元素 key1。
- delete(map1, key1) 删除 map1 中的 key1 元素，不存在 key1 时，不会报错
- map 默认是无序的，不会自动排序
- sort 常用的排序方法
  1. Strings() 字符串排序
  2. Ints() 整数排序
- import . "./pack1" 使用.作为包的别名时，可以省略包名，直接调用其中的变量或方法
- import \_ "./pack1" 使用下划线导入包时，只会执行其中的 init 函数，并初始化全局变量
- 结构体（struct）属于值类型，可以通过 new 函数创建
- new(struct1) 和 &struct1{} 是等价的
- unsafe.Sizeof(T{}) 返回结构体 T 的实例占用的内存大小
- 结构体可以包含一个或多个匿名（或内嵌）字段，结构体可以包含内嵌结构体
- Go 语言中的继承是通过内嵌或组合来实现的
- Go 语言中，一个类型加上它的方法等价于面向对象中的一个类
- 方法的定义必须与接收者类型（receiver_type），在同一个包中被声明
- 类型关联的方法不写在类型结构里面；类型和方法的关联由接收者来建立
- 函数和方法的区别：

  1. 函数将变量作为参数：fn1(recv); 方法在变量上被调用：recv.Method1()
  2. 当接收者或参数为指针时，方法和参数都可以改变其值或状态
  3. 方法中，接收者必须有一个显式的名字，且必须在方法中被使用

- 方法也是可以通过内嵌结构体，实现方法的继承，或者在功能类型中声明一个具名字段
- runtime.GC() 可以显示的触发垃圾回收，回收时会有短暂的性能下降

- 接口定义了一组方法（方法集），不包含方法的实现代码，不能包含变量
- 按约定，只有一个方法的接口名字，由方法名加[e]r 后缀组成
- 接口的类型断言，varI.(\*Struct) 或者 var1.(myInterface)
- Go 语言规范，定义接口方法集的调用规则：

  1. 类型*T 的可调用方法集包含接受者为*T 或 T 的所有方法集
  2. 类型 T 的可调用方法集包含接受者为 T 的所有方法
  3. 类型 T 的可调用方法集不包含接受者为\*T 的方法

- 每个 interface{}变量在内存中占据两个字节，一个用来存储包含它的类型，一个用来存储它包含的数据或者指向数据的指针
- 赋值数据切片至空接口切片，必须使用 for-range 语句一个一个地复制。
- Go 没有类，而是松耦合的类型、方法对接口的实现
- 数据封装，通过标识符的首字母大小写，区分包外和包内可见
- 继承，通过组合实现（内嵌一个以上的类型）
- 多态，通过接口实现（某个类型的实例可以赋给它所实现的任意接口类型的变量）
- panic 一个中止程序的运行时错误。在多层嵌套的函数中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行，并把控制权交给接收到 panic 的函数调用者。这样依次冒泡（执行每层的 defer），直到最顶层，程序崩溃。
- recover 从 panic 或错误场景中恢复
- panic 会导致栈被展开，知道 defer 修饰的 recover()被调用或者程序中止
- go test file_name_test.go --chatty 编译测试程序，--chartty 或 -v 选项，会打印每个执行的测试函数和测试状态
- 用 testing 包中的类型和函数做基准测试时，测试代码必须以 Benchmark 开头的函数，并接收一个\*testing.B 类型的参数
- Go 程序包中必须有的 main 函数也可以看做是一个携程，尽管没有用 go 来启动
- 协程，通过关键字 go 调用一个函数或方法来实现
- 如果在某一时间只有一个协程在执行，不要设置 GOMAXPROCS
- GOMAXPROCS 等同于（并发的）线程数量。
- 通道的声明格式： var identifier chan datatype
- 通道（channel）属于引用类型，可以用 make()函数分配内存，例：int1 := make(chan int)
- 声明带缓冲的通道，例： buf := 100; ch1 := make(chan string, buf)
- 流向通道（发送） 例： ch <- int1
- 通道流出（接收） 例： int2 := <- ch
- 只发送的通道 例： var send_only chan<- int
- 只接收的通道 例： var recv_only <-chan int
- 通道可以显示的关闭，只有发送者需要关闭通道，接收者永远不会需要。
- 使用锁的情景：
  - 访问共享数据结构中的缓存信息
  - 保存应用程序上下文和状态信息数据
- 使用通道的情景：

  - 与异步操作的结果进行交互
  - 分发任务
  - 传递数据所有权

- 信号量模式,一般做法是在 main 函数的最后放置一个{},也可以使用通道让 main 程序等待协程完成。

————————————————

日常开发中，常用的包：

- fmt 包，其中 F 开头的 Print 函数可以直接写入任何 io.Writer，包括文件
  - Println 换行打印
  - Printf 格式化打印，会自动调用 String()方法
  - Errorf 与 Printf 一样，不同的是它会用信息生成一个错误对象，并返回
  - Fprintln 在参数之间添加空格，并追加换行符。
  - Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) 依据指定的格式，向第一个参数内写入字符串
  - Scanln 扫描输入的文本，以空格分隔参数，直到遇到换行
  - Scanf 格式化扫描,第一个参数用作格式字符串，用来决定如何读取
  - Sscan 或 Sscan 开头的函数，从字符串读取，其他与 Scanf 相同。
  - Fscanln
- os 流程包

  - Stdin（standard intput） 标准输入
  - Stdout（standard output） 标准输出
  - File 一个打开文件的描述符，也叫文件句柄
  - Open 打开一个文件，参数是文件名
  - OpenFile 打开一个文件，有三个参数，分别是：
    1. 文件名
    2. 一个或多个标志用|分隔
    3. 使用的文件权限（注意：在读文件的时候，文件的权限是被忽略的，可以用 0 代替，而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。）
  - O_RDONLY 只读
  - O_WRONLY 只写
  - O_RDWR 可读可写
  - O_APPEND 后面追加数据
  - O_CREATE 创建：如果不存在就创建
  - O_EXCL 使用的 O_CREATE，文件必须不存在
  - O_SYNC 打开同步 I/O
  - O_TRUNC 截断：如果文件已存在，就将文件的长度截为 0（全部删除）
  - Exit(v) 终止运行程序

- io 为 I/O 原语提供基本接口。
  - EOF（end of file） 文件的结束符
  - Exit 退出当前程序
- bufio

  - Reader 阅读器为 io 实现缓冲。读者对象。
  - NewReader 返回一个新的阅读器，它的缓冲区具有默认大小。
  - Copy(dst, src) src 拷贝到 dst

- io/ioutil 可以对文件的行读写

  - ReadFile 将整个文件的内容读到一个字符串里，返回一个[]byte
  - WriteFile 将[]byte 的值写入文件，如果没有此文件，会新创建一个

- path/filepath 提供跨平台的函数，处理文件名和路径
  - Base 获得路径中的最后一个元素，不包含后面的分隔符。如：filename := filepath.Base(path)
- compress 提供读取压缩文件的功能，支持的格式：bzip2、flate、gzip、lzw 和 zlib
- flag 解析命令行选项。替换基本常量

  - Parse() 扫描列表并设置 flags
  - Arg(i) 表示第 i 个参数，Parse()之后，全部可用
  - Narg() 返回参数的数量
  - Bool() 定义了一个默认值是 bool 类型的 flag
  - PrintDefaults() 打印 flag 的使用帮助信息

- encoding/json

  - Marshal(v interface{}) ([]byte, error) 编码成 json 结构
  - Unmarshal(data []byte, v interface{}) error 把 json 解码成数据结构
  - MarshalforHTML() 对数据执行 HTML 转码
  - Decoder() 常用 JSON 数据流的读
  - Encoder() 常用 JSON 数据流的写
  - NewEncoder(w io.wWriter) \*Encoder 返回的 Encoder 类型的指针可调用方法 Encode(v interface{})，将数据对象 v 的 json 编码写入 io.Writer（w）中

- net
  - Listen(network, address string) 监听本地网络地址
  - Dial(network, address string) 拨号连接到指定网络上的地址
  - Conn Conn 是一种通用的面向流的网络连接。
- net/http
  - ListenAndServe(address, nill) 监听一个 http 协议的网址
  - ListenAndServeTLS(address, nill) 监听一个 https 协议的网址
  - Redirect(w ResponseWriter, r \*Request, url string, code int) 让浏览器重定向到 url（是请求的 url 的相对路径）以及状态码。
  - NotFound(w ResponseWriter, r \*Request) 这个函数将返回网页没有找到，HTTP 404 错误。
  - Error(w ResponseWriter, error string, code int) 这个函数返回特定的错误信息和 HTTP 代码。
  - HTTP 状态码的常量:
    - http.StatusContinue 100
    - http.StatusOK 200
    - http.StatusFound 302
    - http.StatusBadRequest 400
    - http.StatusUnauthorized 401
    - http.StatusForbidden 403
    - http.StatusNotFound 404
    - http.StatusInternalServerError 500

————————————————

# 断点调试（VS Code）

1. 安装 delve
   windows / linux / OSX  
   go get -u github.com/go-delve/delve/cmd/dlv

2. 设置 launch.json 配置文件
   [参考地址](https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-09-debug.html)

---

优先级 运算符

7 ^ !

6 \* / % << >> & &^

5 + - | ^

4 == != < <= >= >

3 <-

2 &&

1 ||

————————————————

替换标识符,也叫占位符

- %v 以值的默认格式打印，包括数组和结构
- %+v 类似%v,但是输出结构体时会添加字段名
- %#v 值的Go语法表示，实例的完整输出，包括它的字段
- %f float，默认精度打印
- %d 数字，以十进制打印
- %s 字符串，（string类型或[]byte)
- %q 双引号围绕的字符串，由Go语法安全的转义
- %t 布尔
- %p 指针
- %c byte 字节
- %T 打印值的类型
- %% 打印百分号本身

---

未能深刻理解的点：
reflect、method、interface、文件读写、类型转换

# 其他大佬的理解

- [关于 Interface](https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/)
-

---

Go 语言中包含的所属有类型
Bool
Int
Int8
Int16
Int32
Int64
Uint
Uint8
Uint16
Uint32
Uint64
Uintptr
Float32
Float64
Complex64
Complex128
Array
Chan
Func
Interface
Map
Ptr
Slice
String
Struct
UnsafePointer

---

# 出于性能考虑的最佳实践和建议

1. 尽可能使用 := 初始化变量(在函数内部)
2. 尽可能使用字符代替字符串
3. 尽可能使用切片代替数组
4. 尽可能使用数组和切片代替映射
5. 如果只想获取切片中的某项值,不需要值的索引,尽可能的使用for-range遍历切片,这比查询切片中的某个元素要快一些
6. 当数组元素是稀疏的(例如有很多0值或空值nil),使用映射会降低内存消耗
7. 初始化映射时指定其容量
8. 当定义一个方法时,使用指针类型作为方法的接收者
9. 在代码中使用常量或者标志提取常量的值，如 iota
10. 尽可能在需要分配大量内存时使用缓存
11. 使用缓存模板

