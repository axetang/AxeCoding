players = ['charles', 'martina', 'michael', 'florence', 'eli']
print(players[0:3])
print(players[0:len(players)])
print(players[:4])
print(players[3:])
print(players[-3:])

print(players[:])

for player in players[:3]:
    print(player.title())
