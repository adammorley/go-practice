package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	empty = iota
	xVal
	oVal
)
const boardSize = 3

// iota later

// x is represented as 1, o is represented as 2, 0 is not moved yet

type board [boardSize][boardSize]int

// row, column notation
func (b *board) String() string {
	var sb strings.Builder
	for r := 0; r < boardSize; r++ {
		for c := 0; c < boardSize; c++ {
			if b[r][c] == 0 {
				sb.WriteString("-")
			} else if b[r][c] == 1 {
				sb.WriteString("X")
			} else if b[r][c] == 2 {
				sb.WriteString("O")
			}
			if c == 0 || c == 1 {
				sb.WriteString("|")
			} else if c == 2 {
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}

// b is the board, r, c are the locations and v is the value to be stored
func (b *board) add(r, c, v int) {
	if r >= boardSize || c >= boardSize {
		panic("r or c exceeds board size")
	} else if b[r][c] != 0 {
		panic("r, c is already set")
	} else if v != oVal && v != xVal {
		panic("v is invalid")
	}
	b[r][c] = v
}

// note to self: add a generic iterate across board with arbitrary function closure to make this prettier
// is the board fully played?
func (b *board) full() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j] == empty {
				return false
			}
		}
	}
	return true
}

func (b *board) won(v int) bool {
	for r := 0; r < boardSize; r++ {
		if b[r][0] == v && b[r][1] == v && b[r][2] == v{
			return true
		}
	}
	for c := 0; c < boardSize; c++ {
		if b[0][c] == v && b[1][c] == v && b[2][c] == v{
			return true
		}
	}
	if b[0][0] == v && b[1][1] == v && b[2][2] == v {
		return true
	} else if b[0][2] == v && b[1][1] == v && b[2][0] == v {
		return true
	}
	return false
}

// make an arbitrary move
// now i really want a function closure thing here so i can test & set in the board play loop
func (b *board) move() {
	if b.full() {
		panic("board full")
	}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j] == empty {
				b[i][j] = oVal
				return
			}
		}
	}
}

func readInput(prompt string) (v int) {
	for {
		fmt.Println(prompt)
		input, e := bufio.NewReader(os.Stdin).ReadString('\n')
		if e != nil {
			panic("invalid input")
		}
		v, e = strconv.Atoi(strings.Fields(input)[0])
		if e != nil {
			fmt.Println("invalid input, please try again")
		}
		if v < 0 || v >= boardSize {
			fmt.Println("out of bounds")
		} else {
			break
		}
	}
	return
}

func readInputFromUser() (r, c int) {
	fmt.Println("please input numbers one line at a time and press return after each")
	r = readInput("input row number to play")
	c = readInput("input column number to play")
	return
}

func main() {
	var b *board = new(board)
	fmt.Println(b)
	for {
		r, c := readInputFromUser()
		b.add(r, c, xVal)
		if b.won(xVal) {
			fmt.Println("you won!")
            fmt.Println(b)
			break
		}
		b.move()
		if b.won(oVal) {
			fmt.Println("you lost!")
            fmt.Println(b)
			break
		}
		fmt.Println(b)
	}
}

// add error functionality to addToBoard so we can avoid the panic and interpret the error for the human
// iterate makeMove for human and ai
