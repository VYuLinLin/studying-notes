import json

def get_user_name(*arg):
  if len(arg) == 0:
    filename = '../assets/username.json'
  else:
    filename = arg
  try:
    with open(filename) as f_obj:
      username = json.load(f_obj)
  except FileNotFoundError:
    return None
  else:
    return username

def get_new_name():
  """提示用户输入姓名"""
  username = input('What is your name?')
  filename = '../assets/username.json'
  with open(filename, 'w') as f_obj:
    json.dump(username, f_obj)
  return username

def get_user():
  '''问候用户，并指出其名字'''
  username = get_user_name()
  if username:
    print('欢迎回来 ' + str(username) + '!')
  else:
    username = get_new_name()
    print('您的名字已保存到内存中 ' + str(username) + '.')

get_user()