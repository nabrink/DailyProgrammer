def game_of_threes(input):
    if input <= 1:
        return input

    if input % 3 == 0:
        input = input / 3
    else:
        if (input - 1) % == 0:
            input = input - 1
        else:
            input = input + 1
    print(input)
    return game_of_threes(input)


game_of_threes(100)
