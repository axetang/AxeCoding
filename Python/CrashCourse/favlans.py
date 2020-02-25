favlans = {
    'jen': 'python',
    'sarah': 'c',
    'edward': 'ruby',
    'phil': 'python',
}

print("sarah's fav lan is "+favlans['sarah'].title())
print("phil's fav lan is "+favlans['phil'].title())

for name, lan in favlans.items():
    print(name.title()+"'s favorite language is "+lan.title()+".")

for name in favlans.keys():
    print(name.title())

friends = ['phil', 'sarah']
print("---")
for name in favlans.keys():
    print name.title()
    if name in friends:
        print("  Hi "+name.title() +
              ", I see your favorite language is " + favlans[name].title()+"!")


if 'erin' not in favlans.keys():
    print("Erin, please take our poll.")

for name in sorted(favlans.keys()):
    print(name.title()+",thank you for taking the poll.")

print("**********")
for lan in favlans.values():
    print(lan.title())

print("xxxxxxxxx")
for lan in set(favlans.values()):
    print(lan.title())
