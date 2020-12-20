from string import ascii_lowercase


def is_pangram(sentence):
    chars = set(sentence.lower())

    for letter in ascii_lowercase:
        if letter not in chars:
            return False

    return True
