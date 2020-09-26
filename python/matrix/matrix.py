class Matrix:
    def __init__(self, matrix_string):
        rows = matrix_string.splitlines()
        self.matrix = [[int(elm) for elm in r.split(" ")] for r in rows]

    def row(self, index):
        return self.matrix[index - 1]

    def column(self, index):
        return getColumn(index - 1, self.matrix)


def getColumn(index, matrix):
    return [r[index] for r in matrix]
