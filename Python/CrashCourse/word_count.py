def count_words(filename):
    try:
        with open(filename) as f_obj:
            contents = f_obj.read()
    except FileNotFoundError:
        #print("File "+filename+" not exists.")
        pass
    else:
        words = contents.split()
        num_words = len(words)
        print("the file "+filename+" has about "+str(num_words)+" words.")


filename = 'alice.py'
count_words(filename)

filenames = ['alice.py', 'pizza.py',
             'voting.py', 'xxx.py', 'user.py', 'yyy.py']
for filename in filenames:
    count_words(filename)
