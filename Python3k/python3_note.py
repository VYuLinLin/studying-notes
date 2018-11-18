#!/usr/bin/env python
# coding: utf-8

# # Python编程从入门到实践
# 
# # python基础简介
#     - Python 使用Unicode编码
#     - Unicode 一般使用两个字节，ASCII 一般使用一个字节，UTF-8 使用1 至 6个字节
#     - 8bit => 1byte
#     - 2byte => 1字符
# 
# # 数据类型：
#     字符串，整数，浮点数，布尔型，空值，列表[]，元组()，字典{}
#     str, int, float, bool, None, list, tuple, dict
#     range	范围
#     list	是可变的有序集合
#     tuple	是不可变的有序列表
#     dict 	在其他语言中也称为map，是一个键值对的结构
#     set		与dict类似，一组无重复key的集合
#     
# # 字符串方法：
#     ord()		获取字符的整数表示
#     chr()		把编码转换为对应的字符
#     b'ABC'  		Python对bytes类型的数据用带b前缀的单引号或双引号表示
#     encode()		字符编码为指定的字节（bytes）例：'深圳'.encode('utf-8') // b'\xe6\xb7\xb1\xe5\x9c\xb3'
#     decode()		字节解码为指定的字符
#     len()		计算字符串的字符个数 或 列表/元组的长度
#     %		格式化字符串 或 字符串的变量
#             例： 'hello, %s' % 'python' // 'hello, python'
#             '%2d - %02d' % (3,1) // ' 3 - 01'
#             '%.2f' % 3.1415926 // '3.14'
#             字符串中的占位符：
#             %d			整数
#             %f			浮点数
#             %s 			字符串 会把其他类型转换为字符串后在拼接
#             %x  		十六进制整数
#     format()	用{0}占位符格式化字符串
#             例：'Hello,{0}的成绩今年提升了{1:.2f}%'.format('小花',0.66666)
#             // 'Hello,小花的成绩今年提升了0.67%'
#     replace()		全局替换匹配到的字符
#     
# # 列表list方法：
# 	a = []
# 	a[0]				列表通过下标获取值，不能超过索引总值，下标0表示第一个值，-1表示最后一个值，以此类推
# 	a[0:2]			复制索引0到索引2的值，返回新列表
# 	a.append(1)	向列表末尾添加内容，只接收一个参数，少或多都会报错
# 	a.pop()			删除列表末尾的值，并返回这个值，为空时会报错
# 	a.inert(i, val)	在指定位置插入值
# 	a.remove(val)	通过值删除列表元素
# 	del a[i]		通过下标删除列表元素
# 	a.sort(cmp=None, key=None, reverse=False)	无返回值，对原列表进行排序，reverse	 False 升序，True 降序
# 	sorted(iterable[, cmp[, key[, reverse]]])	返回新列表，不会影响原列表，key 和 reverse 比一个等价的 cmp 函数处理速度要快。
# 												这是因为对于每个列表元素，cmp 都会被调用多次，而 key 和 reverse 只被调用一次	
# # dict字典的方法：
# 	d = {key: value}
# 	d[key]			获取或设置对应key的value，无对应key会报错
# 	d.get(key, default)	获取对应key的值，无值时返回None不会报错，并可以指定默认值
# 	d.pop(key)	删除对应的键值对
# 	d.items()		所有的key和value转换成列表
# 	d.keys()		把所有的key转换成列表
# 	d.values()	把所有的value转换成列表
# 	
# 	注：和list比较，dict有以下几个特点：
# 		查找和插入的速度极快，不会随着key的增加而变慢；
# 		需要占用大量的内存，内存浪费多。
# 	而list相反：
# 		查找和插入的时间随着元素的增加而增加；
# 		占用空间小，浪费内存很少。
# 	所以，dict是用空间来换取时间的一种方法。
# 
# 	dict可以用在需要高速查找的很多地方，在Python代码中几乎无处不在，
# 	正确使用dict非常重要，需要牢记的第一条就是dict的key必须是不可变对象。
# 	dict的属性是无序的
#     
# # set方法：
# 	s = set([1,2,3])  ||  s = {1,2,3} 
# 	s.add(key)		添加key到set结构中，重复添加无效果
# 	s.remove(key)	删除key
# 
# -----------------------------------------------------------------------	
# # Python 内置方法:
# 
# ## 类型转换方法：
#     str()	将其他类型转换为字符串类型
#     int()	将字符串、类字节对象、布尔值或数字转换为整数
#     float()	将字符串、类字节对象、布尔值或数字转换为浮点数
#     bool()	将其他类型转换为布尔类型，除0、0.0、''、None、[]、()、{}转换为False，其他值都会转换为True
#     list()	将其他可迭代的数据转换为列表类型
#         例：	
#             list(range(1,4)) // [1,2,3]
#                 list(range(1)) // [0]
#             list((1,)) // [1]
#     tuple()	将可迭代的数据转换为元组类型
#     dict()	将length为1的键值对类型的列表转换为字典类型
#     type()	检查数据属于哪种类型
#     isinstance(value, (类型元组。。。int, str等))	检测参数的是否符合指定的类型集合，返回布尔值
# 
# ## 函数方法：
#     def 方法名()	例：
#         def my_abs(a):
#             if a >= 0:
#                 print(a)
#             else:
#                 pass
# 
# # Python函数总结：
# 	默认参数一定要用不可变对象，如果用的时可变对象，程序运行时会有逻辑错误
# 	*args是可变参数，args接收的是一个tuple；
# 	**kw是关键字参数，kw接收的是一个dict。
# 	定义命名的关键字参数在没有可变参数的情况下不要忘了写分隔符*，否则定义的将是位置参数。
# 	函数执行完毕也没有 return 语句时，自动return None
# 	
# 	abs()		求一个数的绝对值，只能接受一个参数
# 	max(x,y)	求两个数的最大值，可接受任意多个参数
# 	min(x,y)	求两个数的最小值，可接受多个参数
# 
# 	import 模块名/文件名		执行文件并引入文件或模块到当前文件
# 	from 文件名 import 方法名	执行文件并引入文件当中的指定方法
# 
# ## 函数断点方法：
# - (一）命令行中断点：
# 	python -m pdb 文件名
# - (二）脚本中断点：
#     import pdb			引入断点模块
#     pdb.set_trace()		在需要断点的地方调用
# 	- pdb的常用命令说明：
#         l		查看运行到哪行代码
#         n		单步运行，跳过函数
#         s 		单步运行，可进入函数
#         p 		变量 #查看变量值
#         b 		行号 #断点设置到第几行
#         b 		#显示所有断点列表
#         cl 		断点号 #删除某个断点
#         cl 		#删除所有断点
#         c 		#跳到下一个断点
#         r 		#return当前函数
# ------------------------------------------------------------------------------------------
# # 高级特性：
# ## 切片：
# 	a = list(range(100))
# 	a[2:6]		取索引2到索引6的值集合
# 	a[:10]		取前十位数
# 	a[-10:]		取后十位数
# 	a[-1:]		取倒数第一个数
# 	a[-2:-1]	取倒数第二个数
# 	a[:10:2]	前十个数每两位取一个
# 	a[:]		取所有的数
# 	a[::5]		所有的数，每5位取一个
# 	a[::-1]		反转列表集合
# 	元组，字符串与列表同理，不过返回的还是对应的类型
# 
# ## 类（class）
# 	根据类来创建对象被称为实例化
# 
# 	声明一个类并继承另外一个类
# 	class className(object): 			2.7 版本必须定义object,3.0版本可省略
# 		def __init__(self):					初始化方法第一个参数指向实例对象，一般为self
# 			super().__init__()				super代表超级父类，此处是初始化父类的属性或方法
# 
# 	from fileName import className	导入单个类
# 	from fileName import className1,className2	导入多个类
# 	import modelName							导入整个模块
# 	from modelName import *				导入模块中的所有类、方法、变量，不推荐使用此方法导入
# 	from collections import OrderedDict		从标准库 collections 中导入有序字典OrderedDict
# ------------------------------------------------------------------------------------------
# 
# # 文件及错误处理
# 	with open(filePath) as aliasName:
# 		str = aliasName.read() # 字符串形式返回文件文本
# 		lists = aliasName.readlines()  # 读取文件的每一行，返回列表
# 		print(str) 	# 不能直接打印aliasName.read(),可以打印赋值后的变量
# 
# 	with open(filePath, 'w') as fileObj:
# 		fileObj.write('写入的内容')
# 
# 	open方法第二个参数表示打开文件的模式：
# 		r 	只读，也是默认模式，光标在文件开头
# 		r+	读写，不会自动创建文件,如果没有文件，会报fileNotFoundError错误
# 		w	写入，不存在时自动创建，存在时内容会被重置
# 		w+	写读，不存在时自动创建文件
# 		a	附加，不存在时自动创建文件，写的内容会添加后末尾
# 		a+  读写附加，不存在时自动创建文件，写的内容会添加后末尾
#    
# 	open方法会返回一个file对象：
# 		file.write('str')	写入内容
# 		file.read()		读取文件的全部内容
# 		file.readline()		返回一行
# 		file.readlines([size])	返回包含size行的列表，未指定返回所有行
# 
# ------------------------------------------------------------------------------------------
# # 转义字符
# \'　'
# \"	"
# \a	‘bi’响一声
# \b	退格
# \f	换页（在打印时）
# \n	回车，光标在下一行
# \r	换行，光标在上一行
# \t	八个空格
# \\	\
# ---------------------------------------------------------------------------------------
# # Python标准库
# 	collections		容器数据类型集合
# 	random			伪随机数生成器
# 
# # python编码时的注意事项：
# 	文件名不能是关键字
# 	文件命名用驼峰式
# 	类的第一个字母要大写
# 	初始化类时__init__()方法第一个参数尽量用self
# 	python中没有方法提升也没有变量提升
# ---------------------------------------------------------------------------
# # python 一些命令简介：
# 	python --version	查看当前Python的安装版本
# 	pip	--version		查看python的包管理器版本
# 	where python		查看python的安装路径
# 	exit() 		退出
# 	help()		进入帮助文档模式
# 	quit		退出查看帮助模式
# 
# # pip 命令简介：
# 	pip --help 		获取帮助
# 	pip install -U pip	升级pip
# 	pip install SomePackage		安装包
# 	pip uninstall SomePackage	卸载包
# 	pip install -U SomePackage	升级指定的包
# 	pip search SomePackage		搜索包
# 	pip show -f SomePackage		查看指定包的详细信息
# 	pip freeze or pip list 		列出已安装的包
# 	pip list -o 	查看可升级的包
# # jupyter notebook简介
#    
# - Jupyter Notebook 的本质是一个 Web 应用程序，便于创建和共享文学化程序文档，支持实时代码，数学方程，可视化和 markdown。 
#     - 用途包括： 数据清理和转换，数值模拟，统计建模，机器学习等等 
# - Jupyter Notebook 有两种键盘输入模式。
#     - 编辑模式，允许你往单元中键入代码或文本；这时的单元框线是绿色的。
#     - 命令模式，键盘输入运行程序命令；这时的单元框线是灰色。
# - 安装可以使用python自带的pip包管理器安装：pip install jupyter
# - 快捷键
#     - Shift+Enter : 运行本单元，选中下个单元
#     - Ctrl+Enter : 运行本单元
#     - Alt+Enter : 运行本单元，在其下插入新单元
#     - Y：单元转入代码状态
#     - M：单元转入markdown状态
#     - A ：在上方插入新单元
#     - B：在下方插入新单元
#     - X：剪切选中的单元
#     - Shift +V：在上方粘贴单元
# 
# ----------------------------------------------------------------------------

# In[ ]:


print('helo world')


# In[6]:


pwd

