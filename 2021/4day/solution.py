boards = []


with open("input", "r") as file:
    drawn = list(map(int, file.readline()[:-1:].split(",")))
    line = file.readline()
    board = []
    while line:
        if line == "\n" and board:
            boards.append(board)
            board = []
        line = file.readline()
        if line != "\n":
            column_string = line.replace("\n", "")
            column = list(map(int, column_string.split()))
            board.append(column)


def contained(a1, a2):
    for i in a2:
        if i not in a1:
            return False
    return True

d = []

for draw in drawn:
    for board in boards:
        for row in board:
            if contained(d, row):
                t = []
                for i in board:
                    for a in i:
                        if a not in d:
                            t.append(a)
                print(sum(t) * d[-1])
                exit()
    d.append(draw)




for i in boards:
    print(i)
