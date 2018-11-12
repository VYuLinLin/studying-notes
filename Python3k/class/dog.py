# 2.7版本呢需要定义object形参
class Dog():
  # 一次模拟小狗的简单尝试
  def __init__(self, name, age):
    self.name = name
    self.age = age

  def sit(self):
    # 模拟小狗命令式蹲下
    print(self.name.title() + ' is now sitting.')

  def roll_over(self):
    # 模拟小狗命名时打滚
    print(self.name.title() + ' rolled over!')

my_dog = Dog('willie', 6)

print('我的小狗的名字是' + my_dog.name.title())
print('我的小狗今年' + str(my_dog.age) + '岁了')

my_dog.sit()
my_dog.roll_over()