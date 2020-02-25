motorcycles = ['honda', 'yamaha', 'suzuki']
print(motorcycles)
motorcycles[0] = 'ducati'
print(motorcycles)
motorcycles.append('axe')
print(motorcycles)
motorcycles.insert(0, "felix")
motorcycles.insert(2, "mike")
print(motorcycles)
del motorcycles[0]
print(motorcycles)
del motorcycles[1]
del motorcycles[-1]
print(motorcycles)
print("-------------------")
motorcycles.pop()
print(motorcycles)
motorcycles.remove('ducati')
print(motorcycles)
