#from numpy import random

empty = []
lst = [1, 'xiaoming', 29.5, '17312662388']
lst2 = ['001', '2019-11-11', ['三文鱼', '电烤箱']]

print(len(empty), len(lst), len(lst2))

for x in lst:
    print(f'{x}的类型为{type(x)}')

sku = lst2[2]
sku.append('烤鸭')
print(lst2)
sku.insert(1, '牛腱子')
print(lst2)

sku.pop()
print(lst2)
sku.remove(sku[0])
print(lst2)

sku_deep = lst2.copy()
sku_deep[0] = '牛腱'
print(lst2)
print(sku_deep)

a = list(range(1, 20, 3))
print(a)
print(a[:3])
print(a[-1])
print(a[:-1])
print(a[1:5])
print(a[1:5:2])
print(a[::3])
print(a[::-3])


def reverse(lst):
    return lst[::-1]


ra = reverse(a)
print(ra)


a = ()  # 空元组对象
b = (1, 'xiaoming', 29.5, '17312662388')
c = ('001', '2019-11-11', ['三文鱼', '电烤箱'])

'''
# from numpy import random
a = random.randint(1, 5, 10)
print(a)
at = tuple(a)
print(at)
print(at.count(3))
'''
# 可变与不可变
print("可变与不可变:list")
a = [1, 3, [5, 7], 9, 11, 13]
print(a)
a.pop()
a.insert(3, 8)
print(a)
a[2].insert(1, 6)
print(a)

print("可变与不可变:tuple")
a = (1, 2, [5, 7], 9, 11, 13)
a[2].insert(1, 6)
print(a)
