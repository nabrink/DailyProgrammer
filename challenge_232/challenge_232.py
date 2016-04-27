with open("palindrom.txt") as f:
    content = [x.rstrip('\n') for x in f.readlines()]
print(''.join(map(str,content)))
