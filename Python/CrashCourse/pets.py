pets = ['dog', 'cat', 'dog', 'goldfish', 'cat', 'rabbit', 'cat']
print(pets)
while 'cat' in pets:
    pets.remove('cat')

print(pets)


def describe_pet(pet_name, animal_type='dog'):
    print("\nI have a "+animal_type+".")
    print("My "+animal_type+"'s name is "+pet_name.title()+".")


describe_pet("hamster", "harry")
describe_pet("dog", "willie")
describe_pet(pet_name='cles', animal_type='bat')
describe_pet(pet_name='defaulttype')
