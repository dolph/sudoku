package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const unknown = 0b111111111
const zero = 0b000000000
const one = 0b000000001
const two = 0b000000010
const three = 0b000000100
const four = 0b000001000
const five = 0b000010000
const six = 0b000100000
const seven = 0b001000000
const eight = 0b010000000
const nine = 0b100000000

// These are used to convert back and forth between bitwise representations and integers.
var bit2Int = map[int]int{
	zero:  0,
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}
var int2Bit = map[int]int{
	0: zero,
	1: one,
	2: two,
	3: three,
	4: four,
	5: five,
	6: six,
	7: seven,
	8: eight,
	9: nine,
}

type Board struct {
	values [9][9]int
}

func (b Board) PrintChoice(row int, col int, pos int) string {
	if b.values[row][col]&int2Bit[pos] > 0 {
		return strconv.Itoa(pos)
	} else {
		return " "
	}
}

// ╔━━━━━━━╤━━━━━━━╤━━━━━━━╦━
// ║ 1 2 3 | 1 2 3 | 1 2 3 ║
// ║ 4 5 6 | 4 5 6 | 4 5 6 ║
// ║ 7 8 9 | 7 8 9 | 7 8 9 ║
// ╟───────┼───────┼───────╫─
func (b Board) Print() {
	for row := 0; row < 9; row++ {
		if row == 0 {
			fmt.Println("┏━━━━━━━┯━━━━━━━┯━━━━━━━┳━━━━━━━┯━━━━━━━┯━━━━━━━┳━━━━━━━┯━━━━━━━┯━━━━━━━┓")
		} else if row%3 == 0 {
			fmt.Println("┣━━━━━━━┿━━━━━━━┿━━━━━━━╋━━━━━━━┿━━━━━━━┿━━━━━━━╋━━━━━━━┿━━━━━━━┿━━━━━━━┫")
		} else {
			fmt.Println("┠───────┼───────┼───────╂───────┼───────┼───────╂───────┼───────┼───────┨")
		}

		for pos := 0; pos < 3; pos++ {
			for col := 0; col < 9; col++ {
				if col%3 == 0 {
					fmt.Printf("┃")
				} else {
					fmt.Printf("│")
				}
				fmt.Printf(" %s %s %s ", b.PrintChoice(row, col, 3*pos+1), b.PrintChoice(row, col, 3*pos+2), b.PrintChoice(row, col, 3*pos+3))
			}
			fmt.Println("┃")
		}
	}
	fmt.Println("┗━━━━━━━┷━━━━━━━┷━━━━━━━┻━━━━━━━┷━━━━━━━┷━━━━━━━┻━━━━━━━┷━━━━━━━┷━━━━━━━┛")
}

func (b Board) PrintChoices(row int, col int, pass int) {
	if pass == 0 {
		if b.values[row][col]&one > 0 {
			fmt.Printf("1")
		} else {
			fmt.Printf("·")
		}
		if b.values[row][col]&two > 0 {
			fmt.Printf("2")
		} else {
			fmt.Printf("·")
		}
		if b.values[row][col]&three > 0 {
			fmt.Printf("3")
		} else {
			fmt.Printf("·")
		}
	} else if pass == 1 {
		if b.values[row][col]&four > 0 {
			fmt.Printf("4")
		} else {
			fmt.Printf("·")
		}
		if b.values[row][col]&five > 0 {
			fmt.Printf("5")
		} else {
			fmt.Printf("·")
		}
		if b.values[row][col]&six > 0 {
			fmt.Printf("6")
		} else {
			fmt.Printf("·")
		}
	} else if pass == 2 {
		if b.values[row][col]&seven > 0 {
			fmt.Printf("7")
		} else {
			fmt.Printf("·")
		}
		if b.values[row][col]&eight > 0 {
			fmt.Printf("8")
		} else {
			fmt.Printf("·")
		}
		if b.values[row][col]&nine > 0 {
			fmt.Printf("9")
		} else {
			fmt.Printf("·")
		}
	}
}

func initialize() Board {
	rand.Seed(time.Now().Unix())

	b := Board{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// All possible values (representing 1-9) are initialized as valid choices.
			b.values[row][col] = unknown
		}
	}
	return b
}

func main() {
	board := initialize()

	board.Print()
}
