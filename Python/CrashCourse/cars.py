cars = ["bmw", "audi", "toyota", "subaru"]
cars.reverse()
print(cars, sorted(cars))
cars.sort()
print(cars)
cars.sort(reverse=True)
print(cars, len(cars))

for car in cars:
    if car == 'bmw':
        print(car.upper())
    else:
        print(car.title())


car = 'subaru'
print("is car == 'subaru'? I predict true.")
print(car == 'subaru')
