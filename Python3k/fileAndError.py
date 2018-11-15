# 写入文件
def writefile(*str):
  with open('assets/pi_str.txt', 'w') as fileObj:
    print(str)
    if len(str) == 0:
      content = ''
    elif str[0] == 'pi':
      content = '''3.1415926535\n\t8979323846\n\t2643383279'''
    elif str[0] == None:
      content = "print('\a\a\a\a') # 发出连续'bi'的声音"
      print('\a\a\a') # 发出连续'bi'的声音
    fileObj.write(content)

writefile()

# 读取文件
def readfile(path='assets/pi_str.txt'):
  with open(path) as notes:
    note = notes.read()
    # 不能直接打印notes.read(),可以打印赋值后的变量
    print('notes.read() = ',notes.read())
    print('notes.read()变量 = ',note)
    print('note.rstrip() = ', note.rstrip())

