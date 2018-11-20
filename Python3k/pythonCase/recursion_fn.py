# 引入断点包
import pdb

# 递归函数 recursion function

def fact(n):
    if n == 1:
        return 1
    return n * fact(n - 1)
# print(fact(999)) 最大998

# 以上递归会发生栈溢出，解决方法是尾递归
def fact2(n):
    return fact_iter(n, 1)

def fact_iter(num, product):
    if num == 1:
        return product
    return fact_iter(num-1, num * product)

# 汉诺塔游戏
def move(n, a, b, c):
    if n == 1:
        pdb.set_trace()
        print(a, '-->', c)
    else:
        pdb.set_trace()
        move(n-1, a, c, b)
        print(a, '-->', c)
        move(n-1, b, a, c)

'''
函数断点方法：
	（一）命令行中断点：
		python -m pdb 文件名
	（二）脚本中断点：
		import pdb			引入断点模块
		pdb.set_trace()		在需要断点的地方调用
	pdb的常用命令说明：
		l		查看运行到哪行代码
		n		单步运行，跳过函数
		s 		单步运行，可进入函数
　　	p 		变量 #查看变量值
	　　b 		行号 #断点设置到第几行
	　　b 		#显示所有断点列表
	　　cl 		断点号 #删除某个断点
	　　cl 		#删除所有断点
	　　c 		#跳到下一个断点
	　　r 		#return当前函数
	　　exit 	#退出
'''
