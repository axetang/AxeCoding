filename = 'programmig.txt'
with open(filename, 'w') as file_object:
    file_object.write("I love programming!")

with open(filename, 'a') as file_object:
    file_object.write("I love love programming!")
    file_object.write("\nyou love what?")
