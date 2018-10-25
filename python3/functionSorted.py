# -*- coding: utf-8 -*-

L = [('Bob', 75), ('Adam', 92), ('Bart', 66), ('Lisa', 88)]

def by_name(t):
  return t[0].lower()

def by_score(t):
  return t[1]
# 第一种排序方式：有名函数
print(sorted(L, key=by_name))
print(sorted(L, key=by_score, reverse=True))

# 第二种排序方式：匿名函数
print(sorted(L, key=lambda t: t[0].lower()))
print(sorted(L, key=lambda t: t[1], reverse=True))

# 返回函数
# 1法：把外部变量变成容器或者说可变变量
# def createCounter():
#   n = [0]
#   def counter():
#     n[0] = n[0] + 1
#     return n[0]
#   return counter

# 2法：在内部函数里给予外部函数局部变量nonlocal声明，让内部函数去其他领域获取这个变量
def createCounter():
  a = 0
  def counter():
    nonlocal a
    a += 1
    return a
  return counter

# 3法：在内部函数内修改全局变量
# def createCounter():
#   global a
#   a = 0
#   def counter():
#     global a
#     a += 1
#     return a
#   return counter
# 测试:
counterA = createCounter()
print(counterA(), counterA(), counterA(), counterA(), counterA()) # 1 2 3 4 5
counterB = createCounter()
if [counterB(), counterB(), counterB(), counterB()] == [1, 2, 3, 4]:
  print('测试通过!')
else:
  print('测试失败!')