# 高级特性
# 在python中，代码不是越多越好，而是越少越好，不是越复杂越好，而是越简单越好

for a in range(6,10):
    print(a)
    # 6 7 8 9
# 切片
a = list(range(100))
print(a[:10])
print(a[-10])

def trim(str=''):
    str = str_trim(str)
    str = str_trim(str[::-1])
    return str[::-1]

def str_trim(s=''):
    while len(s) != 0 and s[0] == ' ':
        s = s[1:]
    return s