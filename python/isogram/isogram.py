def is_isogram(string):
    found = []

    for letter in string.lower():
        if letter in found:
            return False

        if letter not in (" ", "-"):
            found.append(letter)

    return True
