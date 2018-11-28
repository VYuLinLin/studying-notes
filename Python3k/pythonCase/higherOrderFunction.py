from functools import reduce
# 高阶函数 higher-order function
# 把函数作为参数传入，这样的函数称为高阶函数，函数式编程就是指这种高度抽象的编程范式。


def add(x, y, f):
    return f(x) + f(y)


print(add(-5, 6, abs))  # 11


def f(x):
    return x * x


r = map(f, [0, 1, 2, 3, 4, 5])
print(r)
print(list(r))  # [0, 1, 4, 9, 16, 25]

# 把列表的所有数字转换成字符串
s = list(map(str, [1, 2, 3, 4, 5, 6]))
print(s)

# 序列求和


def adds(x, y):
    return x + y


e = reduce(adds, [1, 2, 3, 4, 5])
print(e)
print(e == sum([1, 2, 3, 4, 5]))  # True

# 首字母大写


def normalize(names):
    return names.title()


p = input('请输入您的姓名：')
n = list(map(normalize, [p]))
print(n)

# 仿float()函数
DIGITS = {'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
          '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
p = input('请输入一个浮点数：')


def strToNum(p):
    def fn(x, y):
        return x * 10 + y

    def charNum(s):
        return DIGITS[s]

    def toNum(n):
        return 0 if len(n) == 0 else reduce(fn, map(charNum, n))
    if len(p.split('.')) == 2:
        a = p.split('.')
        return toNum(a[0]) + (toNum(a[1]) / (10 ** len(a[1])))
    else:
        return toNum(p)


i = strToNum(p)
print(i)
print(isinstance(i, str))
