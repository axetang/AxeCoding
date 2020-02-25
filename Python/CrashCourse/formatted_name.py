def get_formatted_name(firstname, lastname, middle_name=''):
    if middle_name:
        fullname = firstname+' '+middle_name+' ' + lastname
    else:
        fullname = firstname+" "+lastname
    return fullname.title()


musician = get_formatted_name("jimi", "hendrix")
print(musician)
musician = get_formatted_name("jimi", "hendrix", 'lee')
print(musician)
