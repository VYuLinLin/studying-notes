# -*- coding: utf-8 -*-
def triangles(max):
  n, l = 0, [1]
  while n < max:
    yield l
    n = n + 1
    l = [sum(i) for i in zip([0] + l, l + [0])]

def getTriangles(max):
  p = [x for x in triangles(max)]
  return p

t = getTriangles(10)
print(t)

if t == [
  [1],
  [1, 1],
  [1, 2, 1],
  [1, 3, 3, 1],
  [1, 4, 6, 4, 1],
  [1, 5, 10, 10, 5, 1],
  [1, 6, 15, 20, 15, 6, 1],
  [1, 7, 21, 35, 35, 21, 7, 1],
  [1, 8, 28, 56, 70, 56, 28, 8, 1],
  [1, 9, 36, 84, 126, 126, 84, 36, 9, 1]
]:
  print('测试通过！')
else:
  print('测试失败！')

# 迭代器小结

# 凡是可作用于for循环的对象都是Iterable类型；

# 凡是可作用于next()函数的对象都是Iterator类型，它们表示一个惰性计算的序列；

# 集合数据类型如list、dict、str等是Iterable但不是Iterator，不过可以通过iter()函数获得一个Iterator对象。

# Python的for循环本质上就是通过不断调用next()函数实现的，