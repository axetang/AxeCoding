# 判断重复元素
import pyecharts.options as opts
from pyecharts.charts import Scatter
from random import uniform
from random import shuffle
from random import randint, sample


def is_duplicate(lst):
    for x in lst:
        if lst.count(x):
            return True
    return False


a = [1, -2, 3, 4, 1, 2]
b = [x for x in a if a.count(x) == 1]
print("b is ", b)
print(is_duplicate(a))


def reverse_list(lst):
    # 反转列表
    return lst[::-1]


print(reverse_list(a), a)


def find_duplicate(lst):
    # 找出重复元素
    ret = []
    for i in lst:
        if lst.count(i) > 1 and i not in ret:
            ret.append(i)

    return ret


r = find_duplicate([1, 2, 3, 4, 3, 2])
print(r)


def fibonaci(n):
    # 斐波那契数列 1
    if n <= 1:
        return [1]
    fib = [1, 1]
    while len(fib) < n:
        fib.append(fib[len(fib)-1]+fib[len(fib)-2])
    return fib


print("Fibonaci func 1")
f = fibonaci(10)
print(f)


def fibo(n):
    # 斐波那契数列 2
    a, b = 1, 1
    for _ in range(n):
        yield a
        a, b = b, a+b


print("Fibonaci func 2:生成器")
print(list(fibo(2)))
print(list(fibo(5)))


def mode(lst):
    if lst is None or len(lst) == 0:
        return None
    max_frequent_elem = max(lst, key=lambda v: lst.count(v))
    max_frequent = lst.count(max_frequent_elem)
    ret = []
    for i in lst:
        if i not in ret and lst.count(i) == max_frequent:
            ret.append(i)
    return ret


lst = [1, 3, 3, 3, 2, 1, 1, 2]
m = max(lst)
n = max(1, 3, 5)
r = mode(lst)
print(m, n)
print(f'{lst} 中出现次数最多的元素为:{r}')


def max_len(*lists):
    # 带有一个 * 的参数为可变的位置参数，意味着能传入任意多个位置参数。
    # key 函数定义怎么比较大小：lambda 的参数 v 是 lists 中的一个元素
    return max(*lists, key=lambda v: len(v))


r = max_len([1, 2], [2, 4, 5, 7], [9, 0, 1])
print("最长的列表是：", r)


def head(lst):
    # 返回表头 注意返回值中if-else的用法
    return lst[0] if len(lst) > 0 else None


print(head([]))
print(head([1, 2, 3]))


def tail(lst):
    # 返回表尾 注意返回值中if-else的用法
    return lst[-1] if len(lst) > 0 else None


print(tail([]))
print(tail([1, 2, 3]))


def mul_table():
    # 打印乘法表
    print("乘法表")
    for i in range(1, 10):
        for j in range(1, i+1):
            print(str(j)+"*"+str(i)+"="+str(j*i), end='\t')
        print()


mul_table()

# 元素对
# t[:-1]：原列表切掉最后一个元素；
# t[1:]：原列表切掉第一个元素；
# zip(iter1, iter2)：实现 iter1 和 iter2 的对应索引处的元素拼接。
# zip实现列表对应元素处的拼接，如下例
print(list(zip([1, 2], [3, 4])))  # [(1,3),(2,4)]


def pair(t):
    return list(zip(t[:-1], t[1:]))


print(pair([1, 4, 7, 9]))
print(pair(range(10)))


"""
11. 样本抽样
内置 random 模块中，有一个 sample 函数，实现“抽样”功能。
下面例子从 100 个样本中，随机抽样 10 个。
首先，使用列表生成式，创建长度为 100 的列表 lst；
然后，sample 抽样 10 个样本。
"""
#from random import randint,sample
print("样本抽样")
lst = [randint(0, 50) for _ in range(100)]
print(lst[:5])
lst_sample = sample(lst, 10)
print(lst_sample)

"""
12. 重洗数据集
内置 random 中的 shuffle 函数，能冲洗数据。
值得注意，shuffle 是对输入列表就地（in place）洗牌，节省存储空间。
"""
#from random import shuffle
print("重洗数据集")
lst = [randint(0, 50) for _ in range(100)]
shuffle(lst)
print(lst[:5])


#from random import uniform
x, y = [i for i in range(100)], [
    round(uniform(0, 10), 2) for _ in range(100)]
# print(y)


"""
from pyecharts.charts import Scatter
import pyecharts.options as opts
from random import uniform
"""


def draw_uniform_points():
    x, y = [i for i in range(100)], [
        round(uniform(0, 10), 2) for _ in range(100)]
    print(y)
    c = (
        Scatter()
        .add_xaxis(x)
        .add_yaxis('y', y)
    )
    c.render()


draw_uniform_points()
