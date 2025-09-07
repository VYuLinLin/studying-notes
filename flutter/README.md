# flutter + dart 学习笔记
Flutter 是 Google 开发的一个开源 UI 框架，用于从单一的代码库创建精美的、编译型的移动、Web 和桌面应用程序。
学习时的 flutter SDK 3.34.1，dart SDK 3.5.1

学习地址：
[flutter docs](https://docs.flutter.dev/) /
[flutter docs 中文](https://docs.flutter.cn/) /
[flutter sdk api docs](https://api.flutter-io.cn/) /
[flutter&dart基础-掘金](https://juejin.cn/post/7505977477115232319)

flutter 口号： “一切都是(widget)小部件”

Flutter 特点：
  - 跨平台开发：一套代码库可以同时支持 iOS、 Android、web等各种平台。
  - 热重载功能：允许开发者实时查看应用变化，而不需要重新启动应用。
  - 丰富的 UI 组件：提供了丰富的内置组件和丰富的 UI 库。
  - 性能接近原生：由于 Flutter 应用编译为本地 ARM 代码，因此性能接近原生应用。

运行安装环境：
 1. 安装flutter SDK
 2. 安装dart SDK
 3. 安装Android studio
 4. flutter要添加到系统环境变量

基础指令：
- ```dart --version``` 查看dart SDK版本信息
- ```flutter --version``` 查看flutter sdk版本信息
- ```flutter upgrade``` 更新flutter sdk版本到当前渠道的最新版，且更新项目依赖的 packages
- ```flutter channel``` 查看flutter sdk当前渠道
- ```flutter doctor``` 检查flutter运行需要的环境信息
- ```flutter doctor -v``` 检查flutter运行需要的环境详细版本信息
- ```flutter doctor --android-licenses``` 执行并签署SDK许可证
- ```flutter emulators``` 查看可连接的模拟器列表
- ```flutter emulators --launch <emulator id>``` 运行指定id的模拟器
- ```flutter devices``` 查看可连接的设备列表
- ```flutter run``` 运行项目（执行lib\main.dart文件）
- ```flutter clean``` 清除缓存

依赖管理
- ```flutter pub upgrade``` 更新项目依赖到最新的兼容版本
- ```flutter pub get``` 下载项目依赖包
- ```flutter pub outdate``` 显示项目中过时的依赖包

打包构建
- ```flutter build apk/ios``` 打包生成可安装文件

注意事项及重点知识：
  1. flutter默认使用Android studio自带的jdk版本
  2. flutter SDK、JDK、Gradle 三者的版本要兼容匹配才能正常编译
  3. 如果需要指定版本的jdk,直接修改android/gradle.properties文件中的org.gradle.java.home变量的值
  4. 国内使用flutter的镜像地址：FLUTTER_STORAGE_BASE_URL: https://storage.flutter-io.cn     PUB_HOSTED_URL: https://pub.flutter-io.cn
  5. fvm是一个可以管理多个flutter版本的工具，类似fnm\nvm

[JDK对应的Gradle版本](https://docs.gradle.org/current/userguide/compatibility.html#java)

[中国网络环境下使用 Flutter](https://docs.flutter.cn/community/china/)

[Gradle与Gradle Plugin版本匹配说明](https://developer.android.google.cn/build/releases/gradle-plugin?hl=zh-cn)


## dart 语言的特性：
- Dart是面向对象语言，所有对象继承Object，Dart属于强类型语言
- Dart除了可空和基本类型是值传递，其他类型都是引用传递
- Dart是单线程模型，以消息循环机制来运行，包括微任务队列、事件任务队列，其中scheduleMicrotask异步执行微任务，Timer.run或Future异步执行事件任务

## dart 知识点：
- 声明变量
  1. var 固定类型变量
  2. dynamic 动态类型变量
  3. late 延迟初始化变量，首次使用时才初始化
  4. final 只能设置一次的常量
  5. const 编译时常量 或 创建常量值
  6. _  通配符变量，本质是一个占位符，dart sdk 必须大于 3.7
  7. Iterables 迭代方法：for-in、frist、last、firstWhere、singleWhere、any、every、where、takeWhile、skipWhile、map

## Dart 语言内置类型
- 数字 int double，num 包括两种类型
- 字符串 String
- 布尔值 bool
- 记录 (value, value2) 一种匿名、不可变、聚合类型，需dart sdk > 3.0
- 函数 Function
- 列表 List 也叫数组
- 集合 Set
- 映射 Map
- Runes Runes；暴露字符串的 Unicode 码点，也称为Unicode（扩展）字素簇，常被 characters API 替换
- Symbol Symbol
- 空置null Null
其他特殊对象
- Object：除Null外所有Dart类的超类
- Enum: 所有枚举的超类
- Future\Stream: 用于异步编程
- Iterable: 用于for-in 循环和同步生成器函数
- Never： 示表达式永远无法成功完成求值。最常用于总是抛出异常的函数
- dynamic: 表示你希望禁用静态检查。通常你应该改用 Object 或 Object?
- void： 表示一个值永远不会被使用。常用于作为返回类型

## Dart 语言注意事项：
1. 顶层变量和类变量是延迟初始化的；初始化代码在变量首次使用时运行。
2. <font color="red">实例变量可以是 final 但不能是 const</font>
3. 严格来说，级联的“双点”表示法并不是运算符。它只是 Dart 语法的一部分。
4. 模式变量声明必须以 var 或 final 开头，后跟一个模式
5. 混入不能有 extends 或 with 子句，因此 mixin class 也不能有。
6. 类不能有 on 子句，因此 mixin class 也不能有。
7. 所有 Dart 应用都可以使用 async-await、Future 和 Stream 进行非阻塞、交错的计算。然而，<font color="red">Dart Web 平台不支持隔离区（Isolate）</font>。

## 为熟练掌握
- sealed 密封类型

# 出于性能考虑的最佳实践和建议

1. 

