LYRICS = [
    "a Partridge in a Pear Tree",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming",
]

NTH = [
    "first",
    "second",
    "third",
    "fourth",
    "fifth",
    "sixth",
    "seventh",
    "eighth",
    "ninth",
    "tenth",
    "eleventh",
    "twelfth",
]


def recite(start_verse, end_verse):
    return [get_verse(day) for day in range(start_verse - 1, end_verse)]


def get_verse(nth_verse):
    verse = f"On the { NTH[nth_verse] } day of Christmas my true love gave to me: "

    first_gift = LYRICS[0] if nth_verse == 0 else "and " + LYRICS[0]
    gifts = [*LYRICS[nth_verse:0:-1], first_gift]

    verse += ", ".join(gifts)
    return verse + "."
