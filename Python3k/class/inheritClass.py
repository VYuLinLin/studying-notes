# 类的继承

class Car():
  # 一次模拟汽车的属性
  def __init__(self, make, model, year):
    # 模拟汽车的属性
    self.make = make
    self.model = model
    self.year = year
  
  def get_desc_name(self):
    long_name = str(self.year) + ' ' + self.make + ' ' + self.model
    return long_name.title()

class ElectricCar(Car):
  def __init__(self, make, model, year):
    # 初始化父类的属性
    super().__init__(make, model, year)

def person(n):
  print(n)

  
my_tesla = ElectricCar('tesla', 'model s', 2016)
print(my_tesla.get_desc_name())