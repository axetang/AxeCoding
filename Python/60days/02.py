import re
from random import randint, sample
x = 0x18
y = 2.5e6
print(x, y)

lst = [1, 3, 5]
tup = (2, 4, 6)
tup1 = (1,)
# tup2 is a integer,not tuple
tup2 = (1)
print(lst, tup, tup1, tup2)
print(tup[1], tup1[0])
print(type(lst), type(tup1), type(tup2))

dic1 = {1: 3, 2: 5, 3: 7}
dic2 = {'a': 1, 'b': 2, 'c': 3}
set1 = {1, 2, 3, 4, 5}
set2 = {1}
print(type(dic1), type(set1), type(set2))


# 去最求平均
def score_mean(lst):
    lst.sort()
    lst2 = lst[1:-1]
    print(lst2)
    return round((sum(lst2)/len(lst2)), 1)


lst = [2, 9, 8, 6, 3, 7, 4, 1, 5]
mean = score_mean(lst)
print(mean)

# 九九乘法表
print("Multiple Table")
for i in range(1, 10):
    for j in range(1, i+1):
        print(str(j)+'*'+str(i)+'='+str(i*j), end='\t')
    print()

# 随机抽样
print("随机抽样")
lst = [randint(0, 50) for _ in range(100)]
print(len(lst))
print(lst[:5])  # [38, 19, 11, 3, 6]
lst_sample = sample(lst, 10)
print(lst_sample)  # [33, 40, 35, 49, 24, 15, 48, 29, 37, 24]

# 字符串方法
print("字符串方法")
s = ' I love Python '
print(s)
print(s.lstrip())
print(s.rstrip())
print(s.strip())
print(s.replace(" ", "_"))
print(s.find("Python"), s[8])

lst_book = ['this', "is", "a book"]
con = " "
ls = con.join(lst_book)
print(ls)


def is_rotation(s1: str, s2: str) -> bool:
    # 判断 str1 是否由 str2 旋转而来
    if s1 is None or s2 is None:
        return False
    if len(s1) != len(s2):
        return False

    def is_substring(s1: str, s2: str) -> bool:
        return s1 in s2
    return is_substring(s1, s2 + s2)


r = is_rotation('stringbook', 'bookstring')
print(r)  # True
r = is_rotation('greatman', 'maneatgr')
print(r)  # False

# 密码安全测试：正则表达式
pat = re.compile(r'\w{6,20}')
# 这是错误的，因为 \w 通配符匹配的是字母，数字和下划线，题目要求不能含有下划线
# 使用最稳的方法：\da-zA-Z 满足“密码只包含英文字母和数字”
# \d匹配数字 0-9
# a-z 匹配所有小写字符；A-Z 匹配所有大写字符
pat = re.compile(r'[\da-zA-Z]{6,20}')
pat.fullmatch('qaz12')  # 返回 None，长度小于 6
pat.fullmatch('qaz12wsxedcrfvtgb67890942234343434')  # None 长度大于 22
pat.fullmatch('qaz_231')  # None 含有下划线


# 类:使用 __dir__() 查看自带方法：
# 属性前加 2 个 _ 后，变为私有属性,只能在 Dog 类内被共享使用。
# 方法前加 2 个 _ 后，方法变为“私有方法”，只能在 Dog 类内被共享使用。
class Dog(object):
    def __init__(self, name, dtype):
        self.__name = name
        self.__dtype = dtype

    def shout(self):
        print('I\'m %s, type: %s' % (self.name, self.dtype))

    def get_name(self):
        return self.__name


wangwang = Dog('wangwang', 'cute_type')
print(wangwang.__dir__())
# print(wangwang.__name) #AttributeError: 'Dog' object has no attribute '__name'
print(wangwang.get_name())

# 更优雅的做法，使用@property
# 使用 @property 装饰后 name 变为属性，意味着 .name 就会返回这本书的名字，
# 而不是通过 .name() 这种函数调用的方法。这样变为真正的属性后，可读性更好。
# 如果使 name 既可读又可写，就再增加一个装饰器 @name.setter。


class Book(object):
    def __init__(self, name, sale):
        self.__name = name
        self.__sale = sale

    @property
    def name(self):
        return self.__name

    @name.setter
    def name(self, new_name):
        self.__name = new_name

    @property
    def sale(self):
        return self.__sale

    @sale.setter
    def sale(self, new_sale):
        self.__sale = new_sale


a_book = Book('magic_book', 100000)
print(a_book.name, a_book.sale)
a_book.name = 'magic_book_2.0'
a_book.sale = 20000
print(a_book.name, a_book.sale)
