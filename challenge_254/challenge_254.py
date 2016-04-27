def apply_cipher(words):
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    cipher = "zyxwvutsrqponmlkjihgfedcba"

    text = ""
    for word in words:
        text = text + [cipher[alphabet.index(w)] for w, i in word]
