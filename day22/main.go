package main

import (
	"regexp"

	"github.com/ethansaxenian/advent-of-code-2022/util"
)

type point struct {
	r, c, dir int
}

var inputLines = util.FetchInput(22)

func parseInput() ([][]string, []int) {
	move := inputLines[len(inputLines)-1]
	nre := regexp.MustCompile("[0-9]+")
	lre := regexp.MustCompile("L|R")
	numbers := nre.FindAllString(move, -1)
	letters := lre.FindAllString(move, -1)

	moves := []int{}
	for i := 0; i < len(letters); i++ {
		moves = append(moves, util.ToInt(numbers[i]))
		if letters[i] == "R" {
			moves = append(moves, 1)
		} else {
			moves = append(moves, -1)
		}
	}
	moves = append(moves, util.ToInt(numbers[len(numbers)-1]))

	board := [][]string{}
	for _, line := range inputLines[:len(inputLines)-2] {
		row := []string{}
		for i := 0; i < len(inputLines[0]); i++ {
			if i >= len(line) {
				row = append(row, " ")
			} else {
				row = append(row, string(line[i]))
			}
		}

		board = append(board, row)
	}

	return board, moves
}

func getInitialPos(board [][]string) point {
	for r, row := range board {
		for c, p := range row {
			if p == "." {
				return point{r, c, 0}
			}
		}
	}

	return point{-1, -1, -1}
}

func wrap(dir, r, c int, board [][]string) (int, int) {
	switch dir {
	case 1:
		for i := 0; i < len(board); i++ {
			if board[i][c] != " " {
				r = i
				break
			}
		}
	case 3:
		for i := len(board) - 1; i >= 0; i-- {
			if board[i][c] != " " {
				r = i
				break
			}
		}
	case 0:
		for i := 0; i < len(board[0]); i++ {
			if board[r][i] != " " {
				c = i
				break
			}
		}
	case 2:
		for i := len(board[0]) - 1; i >= 0; i-- {
			if board[r][i] != " " {
				c = i
				break
			}
		}
	}
	return r, c
}

func move(pos point, board [][]string) point {
	var dir [2]int
	switch pos.dir {
	case 0:
		dir = [2]int{0, 1}
	case 1:
		dir = [2]int{1, 0}
	case 2:
		dir = [2]int{0, -1}
	case 3:
		dir = [2]int{-1, 0}
	}
	r := pos.r + dir[0]
	c := pos.c + dir[1]
	if r >= len(board) || r < 0 || c >= len(board[0]) || c < 0 || board[r][c] == " " {
		r, c = wrap(pos.dir, r, c, board)
	}
	if board[r][c] == "#" {
		return pos
	}
	return point{r, c, pos.dir}
}

func part1() int {
	board, moves := parseInput()
	pos := getInitialPos(board)
	for j := 0; j < moves[0]; j++ {
		pos = move(pos, board)
	}
	for i := 1; i < len(moves)-1; i += 2 {
		pos.dir = util.Mod((pos.dir + moves[i]), 4)
		for j := 0; j < moves[i+1]; j++ {
			pos = move(pos, board)
		}
	}

	return (pos.r+1)*1000 + (pos.c+1)*4 + pos.dir
}

func moveCube(pos point, board [][]string) point {
	dir := pos.dir
	r := pos.r
	c := pos.c

	var p point
	switch dir {
	case 3:
		if r == 0 {
			if c >= 50 && c < 100 {
				p = point{100 + c, 0, 0}
			} else if c >= 100 && c < 150 {
				p = point{199, c - 100, 3}
			}
		} else if r == 100 {
			if c < 50 {
				p = point{50 + c, 50, 0}
			} else {
				p = point{r - 1, c, 3}
			}
		} else {
			p = point{r - 1, c, 3}
		}

	case 0:
		if c == 149 {
			if r < 50 {
				p = point{(49 - r) + 100, 99, 2}
			}
		} else if c == 99 {
			if r >= 50 && r < 100 {
				p = point{49, 50 + r, 3}
			} else if r >= 100 && r < 150 {
				p = point{(100 - r) + 49, 149, 2}
			} else {
				p = point{r, c + 1, 0}
			}
		} else if c == 49 {
			if r >= 150 && r < 200 {
				p = point{149, r - 100, 3}
			} else {
				p = point{r, c + 1, 0}
			}
		} else {
			p = point{r, c + 1, 0}
		}

	case 1:
		if r == 49 {
			if c >= 100 && c < 150 {
				p = point{c - 50, 99, 2}
			} else {
				p = point{r + 1, c, 1}
			}
		} else if r == 149 {
			if c >= 50 && c < 100 {
				p = point{c + 100, 49, 2}
			} else {
				p = point{r + 1, c, 1}
			}
		} else if r == 199 {
			if c >= 0 && c < 50 {
				p = point{0, c + 100, 1}
			}
		} else {
			p = point{r + 1, c, 1}
		}

	case 2:
		if c == 50 {
			if r >= 0 && r < 50 {
				p = point{(49 - r) + 100, 0, 0}
			} else if r >= 50 && r < 100 {
				p = point{100, r - 50, 1}
			} else {
				p = point{r, c - 1, 2}
			}
		} else if c == 0 {
			if r >= 100 && r < 150 {
				p = point{(100 - r) + 49, 50, 0}
			} else if r >= 150 && r < 200 {
				p = point{0, r - 100, 1}
			}
		} else {
			p = point{r, c - 1, 2}
		}
	}

	if board[p.r][p.c] == "#" {
		return pos
	}

	return p
}

func part2() int {
	board, moves := parseInput()
	pos := getInitialPos(board)

	for j := 0; j < moves[0]; j++ {
		pos = moveCube(pos, board)
	}

	for i := 1; i < len(moves)-1; i += 2 {

		pos.dir = util.Mod((pos.dir + moves[i]), 4)

		for j := 0; j < moves[i+1]; j++ {
			pos = moveCube(pos, board)
		}

	}

	return (pos.r+1)*1000 + (pos.c+1)*4 + pos.dir
}

func main() {
	util.Run(part1, part2)
}
