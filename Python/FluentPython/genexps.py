colors = ['black', 'white']
sizes = ['S', 'M', 'L']
for tshirt in ('%s %s' % (c, s) for c in colors for s in sizes):
    print(tshirt)


l1 = ('egg%s' % i for i in range(10))
print(l1)  # 此时是生成器
print(next(l1))  # 生成器通过next执行一次
print(next(l1))  # 生成器通过next执行一次
print(next(l1))  # 生成器通过next执行一次
print(l1)

coordinates = (33.1, 22.4)
latitude, longitude = coordinates
print(latitude, longitude)

a, b, *rest = range(5)
print(a, b)
print(rest)

a, *body, c, d = range(5)
print(a, body, c, d)

a, b, *cody, d = range(5)
print(a, b, cody, d)

s = 'bicycle'
print(s[::-1])
print(s[::1])
print(s[::2])
print(s[::3])

l = [1, 2, 3]
m = [4, 5]
n = l+m
o = l*3
print(n, o)
ll = [5]
mylist = [ll]*3
print(mylist)
ll[0] = 6
print(mylist)
