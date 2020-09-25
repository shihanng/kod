class Matrix:
    def __init__(self, matrix_string):
        rows = matrix_string.splitlines()
        self.matrix = [[int(elm) for elm in r.split(" ")] for r in rows]
        self.transpose_matrix = list(map(list, zip(*self.matrix)))

    def row(self, index):
        return self.matrix[index - 1]

    def column(self, index):
        return self.transpose_matrix[index - 1]
