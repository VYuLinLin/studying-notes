import math
# 内置函数
a = hex(255)
print(a) # 0xff
type(a)

# n = int(input(''))
# type(n)
# print(n)

# 使用def定义函数
def my_abs(p):
    if not isinstance(p, (int, float)):
        raise TypeError('bad operand type')
    if p >= 0:
        return p
    else:
        return -p

print(my_abs(-5)) # 5

def nop():
    pass

def my_print(s):
    print('我是自定义方法打印的：', s)

def move(x, y, step, angle=0):
    nx = x + step * math.cos(angle)
    ny = y + step * math.sin(angle)
    return nx, ny
# 多次调用会添加多个值
def add_end1(l=[]):
    l.append('END')
    return l
# 多次调用只添加一个值
def add_end2(l=None):
    if l is None:
        l = []
    l.append('None_end')
    return l

# 定义可变参数的函数
def calc(*numbers):
    sum = 0
    for n in numbers:
        sum = sum + n ** 2
    return sum, type(numbers)
    # 调用时：
    # calc(*[1,2,3])  calc(*(1,2,3))    14 tuple
    # calc(*{1,2,3})  calc(*{1,2,3})    14 tuple

# 关键字参数
def person(name, age, **kw):
    print('name:', name, 'age:', age, 'other:', kw)
    print(type(kw)) # <class 'dict'>
    # 调用时：
    # person('victor', 27, **dict)
    # person('victor', 27, gender="女", tel="13256457896")

# 命名关键字参数
def person2(name, age, *, city, job):
    print(name, age, city, job)
    # 使用命名关键字参数时，如果没有可变参数，必须加一个*作为特殊分隔符，
    # 如果缺少*，Python解释器就无法识别位置参数和命名关键字参数

# 参数组合
def f1(a, b=0, *, d, **kw):
    print('必选参数：', a, '默认参数：', b, '命名关键字参数：', d, '关键字参数：', kw)
    # 调用时：
    #  f1(1,d=2)   f1(*('1'), **{'d': 'ddd'})
    # 组合参数定义的顺序必须是：位置参数、默认参数、可变参数、命名关键字参数和关键字参数。
    # 对于任意函数，都可以通过类似func(*args, **kw)的形式调用它

# Python函数总结：
    # 默认参数一定要用不可变对象，如果用的时可变对象，程序运行时会有逻辑错误
    # *args是可变参数，args接收的是一个tuple；
    # **kw是关键字参数，kw接收的是一个dict。
    # 定义命名的关键字参数在没有可变参数的情况下不要忘了写分隔符*，否则定义的将是位置参数。
    