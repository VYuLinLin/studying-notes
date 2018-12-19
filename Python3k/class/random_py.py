# random(伪随机数生成器)是python的标准库之一
# randint 直接生成整数的随机数
from random import randint

x = randint(1, 6)    # 生成一个1至6范围内的整数

i = 10
while i > 0:
  i -=1
  print('randint生成的随机数：' + str(randint(1, 6)))