# Python的两种循环
# for in 循环
# a = 0
# for x in list(range(101)):
#     print(type(x))
#     a = a + x
# print(a) // 5050

# while 循环
sum = 0
n = 9
while n > 0:
    sum = sum + n
    n = n - 1
    print(sum)
print('while循环最后的值：', sum)  # while循环最后的值： 45

# break 提前退出循环,后面的代码不会再执行
n = 1
while n <= 100:
    if n > 10:
        break
    print(n)
    n = n + 1
print('break提前结束循环', 'END')

# continue 退出当次循环,当次后面的代码不会再执行
n = 1
while n < 10:
    n = n + 1
    if n % 2 == 0:
        continue
    print(n)
print('continue提前结束循环', 'END')

# 有些时候，如果代码写得有问题，会让程序陷入“死循环”，也就是永远循环下去。
# 这时可以用Ctrl+C退出程序，或者强制结束Python进程。