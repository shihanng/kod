def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("sequences have no equal length")

    hamming_distance = 0

    for i, elem_a in enumerate(strand_a):
        if elem_a != strand_b[i]:
            hamming_distance += 1

    return hamming_distance
