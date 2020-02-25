# 判断重复元素
def is_duplicate(lst):
    for x in lst:
        if lst.count(x):
            return True
    return False


a = [1, -2, 3, 4, 1, 2]
print(is_duplicate(a))

# 反转列表


def reverse_list(lst):
    return lst[::-1]


print(reverse_list(a), a)

'''找出重复元素'''


def find_duplicate(lst):
    ret = []
    for i in lst:
        if lst.count(i) > 1 and i not in ret:
            ret.append(i)

    return ret


r = find_duplicate([1, 2, 3, 4, 3, 2])
print(r)

# 斐波那契数列


def fibonaci(n):
    if n <= 1:
        return [1]
    fib = [1, 1]
    while len(fib) < n:
        fib.append(fib[len(fib)-1]+fib[len(fib)-2])
    return fib


f = fibonaci(10)
print(f)


def fibo(n):
    a, b = 1, 1
    for _ in range(n):
        yield a
        a, b = b, a+b


print(list(fibo(2)))
print(list(fibo(5)))


def mode(lst):
    if lst is None or len(lst) == 0:
        return None
    return max(lst, key=lambda v: lst.count(v))


lst = [1, 3, 3, 2, 1, 1, 2]
r = mode(lst)
print(f'{lst} 中出现次数最多的元素为:{r}')


# 返回表头
def head(lst):
    return lst[0] if len(lst) > 0 else None


print(head([]))
print(head([1, 2, 3]))

# 返回表尾


def tail(lst):
    return lst[-1] if len(lst) > 0 else None


print(tail([]))
print(tail([1, 2, 3]))

# 打印乘法表


def mul_table():
    for i in range(1, 10):
        for j in range(1, i+1):
            print(str(j)+"*"+str(i)+"="+str(j*i), end='\t')
        print()


mul_table()
