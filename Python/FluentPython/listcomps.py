x = 'ABC'
dummy = [ord(x) for x in x]
print(x)
print(dummy)

symbols = '$%^&*()9'
beyond_ascii = [ord(s) for s in symbols if ord(s) > 17]
print(beyond_ascii)

print("笛卡尔积")
colors = ['black', 'white', 'red']
sizes = ['S', 'M', 'L']
tshirts = [(color, size) for color in colors for size in sizes]
print(tshirts)

vec = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
nums = [num for elem in vec for num in elem]
print(nums)
