
def parse_input():
    draws = []
    boards = []

    with open("day4_input.txt", "r") as file:
        draws.extend([int(i) for i in file.readline().split(",")])

        newboard = [[], [], [], [], []]
        i = 0
        for line in file:
            if line.strip():
                newboard[i].extend([int(x) for x in line.strip().split(" ") if x])
                i += 1
            else:
                boards.append(newboard)
                newboard = [[], [], [], [], []]
                i = 0

        boards.append(newboard)

    return draws, [b for b in boards if b[0]]


def has_bingo(board):
    if any(all(item == "X" for item in row) for row in board):
        return True
    for i in range(5):
        if all(row[i] == "X" for row in board):
            return True
    return False


def mark_draw(board, draw):
    for x in range(5):
        for y in range(5):
            if board[x][y] == draw:
                board[x][y] = "X"


def print_board(board):
    for row in board:
        for i in row:
            print(i, end=" ")
        print("\n")


def score_board(board):
    s = 0
    for x in range(5):
        for y in range(5):
            if board[x][y] != "X":
                s += board[x][y]
    return s


if __name__ == '__main__':
    draws, boards = parse_input()

    for n in draws:
        for board in boards:
            mark_draw(board, n)
            if has_bingo(board):
                print(n)
                print_board(board)
                print(n * score_board(board))
                exit(0)
