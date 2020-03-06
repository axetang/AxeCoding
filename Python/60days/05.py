# 字典（dict），一种映射对象（mapping）类型，键值对的容器。五种创建方法
# 创建 dict 的方法很简单，使用一对中括号 {}，键值对中间用冒号，每项间使用逗号分割。
empty = {}
dic = {'a': 1, 'c': 3, 'e': 5}
dic1 = dict(a=1, b=2, c=3)
dic2 = dict({'a': 1, 'b': 2}, c=3, d=4)
print(dic, dic1, dic2)
dic3 = dict([('a', 1), ('b', 2)], c=5, d=7)
print(dic3)

dic4 = {}.fromkeys(['a', 'b', 'c'], [1, 2, 3])
print('dic4\n', dic4)
dic5 = {'a': 1, 'b': 2}.fromkeys(['c', 'd'], [3, 4])
print('dic5\n', dic5)

# 基本操作
print("Dict 基本操作：")
d = {'a': 1, 'b': 2, 'c': 3}
for key, value in d.items():
    print(key, value)
keys = set(d)
print(keys)
keys = set(d.keys())
print(keys)
values = set(d.values())
print(values)

if 'c' in d:
    print('key c is in d')
if 'e' not in d:
    print('key e is not in d')
if 'c' in d.keys():
    print('key c is in d')
if 'e' not in d.keys():
    print('key e is not in d')

print("key c's value:", d.get('c'))
print(d)
d['d'] = 4
print(d)
del d['d']
print(d)
d.pop('c')
print(d)
d['e'] = 10
print(d)
print(d.keys())
print(d.values())
print(d.items())

'''
可哈希的对象才能作为字典的键，不可哈希的对象不能作为字典的键。
因为列表是可变对象，而可变对象是不可哈希的，所以会抛出如上异常。
l = [1, 2]
dic = {l: 1}
print(dic)
'''
# 集合:集合是一种不允许元素出现重复的容器。


def duplicated(lst):
    return len(lst) != len(set(lst))


print(duplicated([1, 3, 4, 6, 2, 1, 3]))

# 与字典（dict）类似，集合（set）也是由一对花括号（{}）创建。
# 但是，容器内的元素不是键值对。
# 同字典类似，集合内的元素必须是可哈希类型（hashable）。
# 这就意味着 list、dict 等不可哈希的对象不能作为集合的元素。
a = {1, 2, 3}
# 另一种创建集合的方法，是通过 Python 的内置的 set 函数，
# 参数类型为可迭代对象 Iterable。
a = set([1, 3, 5, 6, 3])
print(a)

# 集合方法
a, b, c, e = {1, 3, 5, 7}, {3, 4, 5, 6}, {6, 7, 8, 9, 3}, {1, 3, 5, 7, 9, 0}
d = a.union(b, c)
print(d)

print(a.difference(b, c))
print(a.intersection(b, c))
print(a.issubset(b))
print(a.issubset(a))
print(a.issubset(e))
