def to_palindromic(number, step):
    if len(number) <= 1 or step > 1000 or is_palindromic(number):
        return number
    else:
        return to_palindromic(str(int(number) + int(number[::-1])), step + 1)

def is_palindromic(number):
    return number == number[::-1]

number = input("Enter a number: ")
print(to_palindromic(number, 0))
