alien = {"color": "green", "points": 5}
print(alien)
print(alien['color'], alien['points'])
alien['x'] = 0
alien['y'] = 25
print(alien)

alien0 = {}
alien0['color'] = 'green'
alien0['points'] = 5
print(alien0)
alien0['color'] = 'yellow'
print(alien0)

#del alien0['points']
print(alien0, "\n****************************")


al0 = {'color': 'green', "points": 5}
al1 = {'color': 'yello', "points": 10}
al2 = {'color': 'red', "points": 15}
als = [al0, al1, al2]
for al in als:
    print(al)
print("\n------------------------\n")

als = []
for alnum in range(30):
    newal = {"color": "green", 'points': 5, "speed": 'slow'}
    als.append(newal)

for al in als[:5]:
    print(al)
print("total number of als: "+str(len(als)))
