# 需求：
# 有一千瓶饮料，三个空瓶可以换一瓶饮料，不得借瓶，问可以喝多少瓶饮料？
# 思路 
# 利用递归方法累计饮料个数，
# 饮料喝完需要三个空瓶换一瓶饮料，
# 所以需要循环除以3，直至不能被3整除，
# 最后合计数在加上不能被3整除的饮料个数，
# 就是合计可以喝的饮料个数。

# //（地板除，就是只取余数的整数部分）
def g(n=1000):
    if n < 3:
        return n
    m = n // 3 + (n % 3)
    return n + g(m) - (n % 3)

print(g(1000))