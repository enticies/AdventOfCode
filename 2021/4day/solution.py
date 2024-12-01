def main():
    boards = []

    with open("input", "r") as file:
        draws = list(map(int, file.readline()[:-1:].split(",")))
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

    return boards, draws


def get_cols(array):
    cols = []

    for i in range(len(array[0])):
        t = []
        for a in array:
            t.append(a[i])
        cols.append(t)

    return cols

def check_board(board, draws):
    cols = get_cols(board)
    
    for col in cols:
        if contained(draws, col):
            return True

    for row in board:
        if contained(draws, row):
            return True
    return False
    
def contained(a1, a2):
    for i in a2:
        if i not in a1:
            return False
    return True

def calculate_board(board, draws):
    t = 0
    for row in board:
        for item in row:
            if item not in draws:
                t += item

    return t * draws[-1]

def find_winner(boards, draws):
    d = []
    for draw in draws:
        for board in boards:
            if check_board(board, d):
                return calculate_board(board, d)
        d.append(draw)
    return None

def find_last_winner(boards, draws):
    scores = []
    d = []

    for draw in draws:
        for board in boards:
            if check_board(board, d):
                winner = board
                scores.append(calculate_board(board, d))
                boards.remove(board)
        d.append(draw)

    return scores 


if __name__ == "__main__":
    boards, draws = main()

    # score = find_winner(boards, draws)
    score = find_last_winner(boards, draws)
    print(score)

