def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("sequences have no equal length")

    hamming_distance = 0

    for elem_a, elem_b in zip(strand_a, strand_b):
        if elem_a != elem_b:
            hamming_distance += 1

    return hamming_distance
