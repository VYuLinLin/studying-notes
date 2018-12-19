# collections(容器数据类型集合)是python的标准库之一
# OrderedDict 的作用是记住字典的添加顺序
from collections import OrderedDict

ul = dict()

ul['a'] = 'A'
ul['b'] = 'B'
ul['c'] = 'C'

ol = OrderedDict()

ol['c'] = 'C'
ol['b'] = 'B'
ol['a'] = 'A'



def printDict(d):
  for name,language in d.items():
    print(name.title() + '---------------' + language.title())

print('dict数据结构：')
printDict(ul)
print('\nOrderedDict数据结构：')
printDict(ol)
print('\nul是否与ol相等：')
print(ul == ol)   # True
'''
  在Python 3.6之前，常规的dict不会跟踪插入顺序，
  在它上面进行迭代会根据键在哈希表中的存储方式产生值，而哈希表又会受到随机值的影响，以减少冲突。
  相反，在OrderedDict中，在创建迭代器时记住并使用插入项的顺序。
  在Python 3.6及以上版本中，内置的dict会跟踪插入顺序，尽管这种行为是实现更改的副作用，不应该依赖。
'''