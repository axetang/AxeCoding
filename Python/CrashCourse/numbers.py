for value in range(1, 5):
    print(value)

numbers = list(range(1, 6))
print(numbers)

evennumbers = list(range(2, 11, 2))
print(evennumbers)

squares = []
for value in range(1, 11):
    square = value**2
    squares.append(square)

print(squares)

print(min(squares), max(squares), sum(squares))

sq = [val**2 for val in range(1, 11)]
print(sq)
