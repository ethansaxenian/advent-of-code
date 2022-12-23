from day4a import parse_input, mark_draw, has_bingo, print_board, score_board

draws, boards = parse_input()

for n in draws:
    for board in boards:
        mark_draw(board, n)
    if len(boards) == 1:
        if has_bingo(boards[0]):
            print_board(boards[0])
            print(n * score_board(boards[0]))
            exit(0)
    boards = [b for b in boards if not has_bingo(b)]
