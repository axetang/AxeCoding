# 1. update
# 实际使用字典时，需要批量插入键值对到已有字典中，
# 使用 update 方法实现批量插入。已有字典中批量插入键值对：
d = {'a': 1, 'b': 2}
d.update({'c': 3, 'd': 4, 'e': 5})
print(d)
d.update([('f', 6), ('g', 7)])
print(d)
d.update([('h', 8), ('i', 9)], j=10)
print(d)

# 2. setdefault
# 如果仅当字典中不存在某个键值对时，才插入到字典中；
# 如果存在，不必插入（也就不会修改键值对）。
# 这种场景，使用字典自带方法 setdefault
d = {'a': 1, 'b': 2}
d.setdefault('c', 3)
print(d)
d.setdefault('c', 4)
print(d)


# 3. 字典并集
# 先来看这个函数 f，为了好理解，显示的给出参数类型、返回值类型，这不是必须的。
print("*/**")


def f(d: dict) -> dict:
    return {**d}


def merge(d1, d2):
    return {**d1, **d2}


def merge1star(d1, *d2):
    return [d1, *d2]


def merge2star(d1, **d2):
    return {**d1, **d2}


print(f({'a': 1, 'b': 2}))
print(merge({'a': 1, 'b': 2}, {'c': 3}))
print(merge1star(5, 6, 7, 8))
print(merge2star({'a': 1, 'b': 2}, c=3, d=4))
